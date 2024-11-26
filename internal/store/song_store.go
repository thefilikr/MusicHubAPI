package store

import (
	"database/sql"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
)

type Song struct {
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate *string   `json:"release_date,omitempty"`
	Text        *[]string `json:"text,omitempty"`
	Link        *string   `json:"link,omitempty"`
}

type SongStore interface {
	CreateSong(group, song string, releaseDate, link *string, text *[]string) error
	GetSong(group, songName, countVerses, numPage string) (*Song, error)
	GetSongs(group, songName, releaseDate, link *string, text *[]string) (*[]Song, error)
	GetIDSong(group, songName string) (*uint, error)
	DeleteSong(id uint) error
	EditSong(id uint, group, song string, releaseDate, link *string, text *[]string) error
}

type SQLSongStore struct {
	DB  *sql.DB
	Log *slog.Logger
}

func (s *SQLSongStore) CreateSong(group, songName string, releaseDate, link *string, text *[]string) error {
	query := `
		INSERT INTO songs (group_name, song, release_date, link) 
		VALUES ($1, $2, 
			CASE WHEN $3 = '' THEN NULL ELSE $3 END, 
			CASE WHEN $4 = '' THEN NULL ELSE $4 END
		)
		RETURNING id_song`

	var id int
	err := s.DB.QueryRow(query, group, songName, releaseDate, link).Scan(&id)

	if err != nil {
		s.Log.Error("Failed create song err: %w", err.Error())
		return err
	}

	s.Log.Debug(fmt.Sprintf("Inserted song with ID: %d", id))

	if text == nil {
		return err
	}

	for i, verse := range *text {
		query := `
			INSERT INTO verses (text_verses, song_id, verses_num)
			VALUES ($1, $2, $3)`
		_, err := s.DB.Exec(query, verse, int(id), i+1)
		if err != nil {
			s.Log.Error(err.Error())
			return err
		}
	}

	return err
}

func (s *SQLSongStore) GetSong(group, songName, countVerses, numPage string) (*Song, error) {
	query := `SELECT id_song, group_name, song, release_date, link FROM songs WHERE group_name = $1 AND song = $2`
	row := s.DB.QueryRow(query, group, songName)

	var song Song
	var id uint
	if err := row.Scan(
		&id,
		&song.Group,
		&song.Song,
		&song.ReleaseDate,
		&song.Link,
	); err != nil {
		s.Log.Error(err.Error())
		if err == sql.ErrNoRows {
			s.Log.Error("Failed get song", sql.ErrNoRows.Error())
			return nil, nil
		}
		return nil, err
	}

	s.Log.Debug("Successful get song")

	var err error
	song.Text, err = s.GetPaginationVerses(id, countVerses, numPage)
	if err != nil {
		s.Log.Debug("Failed get verses: ", err.Error())
		return &song, err
	}

	s.Log.Debug("Successful get verses")

	return &song, nil
}

func (s *SQLSongStore) GetSongs(group, songName, releaseDate, link *string, text *[]string) (*[]Song, error) {

	var songs []Song

	query := `SELECT DISTINCT s.id_song FROM songs s LEFT JOIN verses v ON s.id_song = v.song_id WHERE `
	var args []interface{}
	var updates []string

	if group != nil {
		updates = append(updates, `s.group_name = $`+fmt.Sprint(len(args)+1))
		args = append(args, *group)
	}
	if songName != nil {
		updates = append(updates, `s.song = $`+fmt.Sprint(len(args)+1))
		args = append(args, *songName)
	}
	if releaseDate != nil {
		updates = append(updates, `s.release_date = $`+fmt.Sprint(len(args)+1))
		args = append(args, *releaseDate)
	}
	if link != nil {
		updates = append(updates, `s.link = $`+fmt.Sprint(len(args)+1))
		args = append(args, *link)
	}

	if text != nil {
		for _, verse := range *text {
			updates = append(updates, "v.text_verses LIKE '%' || $"+fmt.Sprint(len(args)+1)+" || '%' ")
			args = append(args, verse)
		}
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("not data")
	}

	query += strings.Join(updates, " OR ")

	s.Log.Debug("Quere: ", query)
	s.Log.Debug("Updates: ", updates)
	s.Log.Debug("Args: ", args)

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error when executing the request: %w", err)
	}

	var ids []uint
	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("error reading the string: %w", err)
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error when traversing lines: %w", err)
	}

	rows.Close()

	if len(ids) == 0 {
		return nil, nil
	}

	query = `SELECT song, group_name, release_date, link FROM songs WHERE id_song = $1`
	for _, id := range ids {
		row := s.DB.QueryRow(query, id)

		var song Song
		if err := row.Scan(
			&song.Song,
			&song.Group,
			&song.ReleaseDate,
			&song.Link,
		); err != nil {
			s.Log.Error(err.Error())
			if err == sql.ErrNoRows {
				s.Log.Error("Failed get song", sql.ErrNoRows.Error())
				return &songs, nil
			}
			return nil, err
		}

		s.Log.Debug("Successful get song")

		verses, err := s.GetVerses(id)
		if err != nil {
			s.Log.Debug("Failed get verses: ", err.Error())
			return &songs, err
		}

		song.Text = &verses
		s.Log.Debug("Successful get verses")

		songs = append(songs, song)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error when traversing lines: %w", err)
	}

	return &songs, nil
}

func (s *SQLSongStore) GetIDSong(group, songName string) (*uint, error) {
	query := `SELECT id_song FROM songs WHERE group_name = $1 AND song = $2`
	row := s.DB.QueryRow(query, group, songName)

	var id uint
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			s.Log.Debug("Error get id song: ", sql.ErrNoRows)
			return nil, nil
		}
		s.Log.Debug("Error get id song: ", err.Error())
		return nil, err
	}

	return &id, nil
}

func (s *SQLSongStore) DeleteSong(id uint) error {
	query := `DELETE FROM songs WHERE id_song = $1`

	_, err := s.DB.Exec(query, id)
	return err
}

func (s *SQLSongStore) EditSong(id uint, group, songName string, releaseDate, link *string, text *[]string) error {
	query := `UPDATE songs SET `
	var args []interface{}
	var updates []string

	if group != "" {
		updates = append(updates, `group_name = $`+fmt.Sprint(len(args)+1))
		args = append(args, group)
	}
	if songName != "" {
		updates = append(updates, `song = $`+fmt.Sprint(len(args)+1))
		args = append(args, songName)
	}
	if releaseDate != nil {
		updates = append(updates, `release_date = $`+fmt.Sprint(len(args)+1))
		args = append(args, *releaseDate)
	}
	if link != nil {
		updates = append(updates, `link = $`+fmt.Sprint(len(args)+1))
		args = append(args, *link)
	}

	if len(updates) == 0 {
		return fmt.Errorf("not data")
	}

	query += strings.Join(updates, ", ") + ` WHERE id_song = $` + fmt.Sprint(len(args)+1)
	args = append(args, id)

	s.Log.Debug("Quere: ", query)
	s.Log.Debug("Args: ", args)

	_, err := s.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	s.Log.Debug("Successful edit args song")

	if text == nil {
		return nil
	}

	query = `DELETE FROM verses WHERE song_id = $1`

	if _, err := s.DB.Exec(query, id); err != nil {
		s.Log.Debug("Failed delete verses")
		return err
	}

	s.Log.Debug("Successful delete verses")

	for i, verse := range *text {
		query := `
			INSERT INTO verses (text_verses, song_id, verses_num)
			VALUES ($1, $2, $3)`
		_, err := s.DB.Exec(query, verse, id, i+1)
		if err != nil {
			return err
		}
	}

	s.Log.Debug("Successful edit verses")

	return err
}

func (s *SQLSongStore) GetVerses(id uint) ([]string, error) {
	query := `SELECT text_verses FROM verses WHERE song_id = $1 ORDER BY verses_num ASC`
	rows, err := s.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error when executing the request: %w", err)
	}
	defer rows.Close()

	var verses []string

	for rows.Next() {
		var verse string
		if err := rows.Scan(&verse); err != nil {
			return nil, fmt.Errorf("error reading the string: %w", err)
		}

		verses = append(verses, verse)
	}

	// Проверяем на ошибки после обхода строк
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error when traversing lines: %w", err)
	}

	return verses, nil
}

func (s *SQLSongStore) GetPaginationVerses(id uint, countVerses, numPage string) (*[]string, error) {

	var verses []string
	if countVerses == "" || numPage == "" {
		var err error
		verses, err = s.GetVerses(id)
		if err != nil {
			s.Log.Error("Conversion error:", err)
			return nil, err
		}
		return &verses, nil
	}

	numCountVerses, err := strconv.ParseUint(countVerses, 10, 0)
	if err != nil {
		s.Log.Error("Conversion error:", err)
		return &verses, nil
	}
	uintCountVerses := uint(numCountVerses)

	numNumPage, err := strconv.ParseUint(numPage, 10, 0)
	if err != nil {
		s.Log.Error("Conversion error:", err)
		return &verses, nil
	}
	uintNumPage := uint(numNumPage)

	totalPages := math.Ceil(float64(len(verses)) / float64(uintCountVerses))

	if uintNumPage > uint(totalPages) {
		return nil, nil
	}

	verses = verses[:0]

	query := `SELECT text_verses FROM verses WHERE song_id = $1 ORDER BY verses_num ASC LIMIT $2 OFFSET $3`

	rows, err := s.DB.Query(query, id, uintCountVerses, uintNumPage*uintCountVerses)
	if err != nil {
		return nil, fmt.Errorf("error when executing the request: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var verse string
		if err := rows.Scan(&verse); err != nil {
			return nil, fmt.Errorf("error reading the string: %w", err)
		}
		verses = append(verses, verse)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error when traversing lines: %w", err)
	}

	return &verses, nil
}
