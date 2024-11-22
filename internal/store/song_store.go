package store

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Song struct {
	ID          int        `json:"id"`
	Group       string     `json:"group"`
	Song        string     `json:"song"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Text        *string    `json:"text,omitempty"`
	Link        *string    `json:"link,omitempty"`
}

type SongStore interface {
	CreateSong(group, song string, releaseDate *time.Time, text, link *string) error
	GetSong(group, songName string) (*Song, *uint, error)
	DeleteSong(id uint) error
	EditSong(id uint, group, song string, releaseDate *time.Time, text, link *string) error
}

type SQLSongStore struct {
	DB *sql.DB
}

func (s *SQLSongStore) SaveSong(group, songName string, releaseDate *time.Time, text, link *string) error {
	query := `
		INSERT INTO tasks (group, song, release_date, text, link) 
		VALUES ($1, $2, 
			CASE WHEN $3 = '' THEN NULL ELSE $3 END, 
			CASE WHEN $4 = '' THEN NULL ELSE $4 END, 
			CASE WHEN $5 = '' THEN NULL ELSE $5 END
		)`
	_, err := s.DB.Exec(query, group, songName, releaseDate, text, link)
	return err
}

func (s *SQLSongStore) GetSong(group, songName string) (*Song, *uint, error) {
	query := `SELECT id_song, group, song, release_date, text, link FROM tasks WHERE group = $1, song = $2`
	row := s.DB.QueryRow(query, group, songName)

	var song Song
	var id uint
	err := row.Scan(
		id,
		song.Group,
		song.Song,
		&song.ReleaseDate,
		&song.Text,
		&song.Link,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	return &song, &id, nil
}

func (s *SQLSongStore) DeleteSong(id uint) error {
	query := `DELETE tasks WHERE id_song = $1`

	_, err := s.DB.Exec(query, id)
	return err
}

func (s *SQLSongStore) EditSong(id uint, group, songName string, releaseDate *time.Time, text, link *string) error {
	query := `UPDATE tasks SET `
	var args []interface{}
	var updates []string

	if group != "" {
		updates = append(updates, `group = $`+fmt.Sprint(len(args)+1))
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
	if text != nil {
		updates = append(updates, `text = $`+fmt.Sprint(len(args)+1))
		args = append(args, *text)
	}
	if link != nil {
		updates = append(updates, `link = $`+fmt.Sprint(len(args)+1))
		args = append(args, *link)
	}

	if len(updates) == 0 {
		return fmt.Errorf("not data")
		// TODO изменить реализацию логирования
	}

	query += strings.Join(updates, ", ") + ` WHERE id = $` + fmt.Sprint(len(args)+1)
	args = append(args, id)

	_, err := s.DB.Exec(query, args...)
	return err
}
