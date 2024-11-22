package store

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Song struct {
	ID          int
	Group       string
	Song        string
	ReleaseDate *time.Time
	Text        *string
	Link        *string
}

type SongStore interface {
	createSong(group, song string, releaseDate *time.Time, text, link *string) error
	getSong(group, songName string) (*Song, *uint, error)
	deleteSong(id uint) error
	editSong(id int, group, song string, releaseDate *time.Time, text, link *string) error
}

type SQLSongStore struct {
	DB *sql.DB
}

func (s *SQLSongStore) createSong(group, song string, releaseDate *time.Time, text, link *string) error {
	query := `
		INSERT INTO tasks (group, song, release_date, text, link) 
		VALUES ($1, $2, 
			CASE WHEN $3 = '' THEN NULL ELSE $3 END, 
			CASE WHEN $4 = '' THEN NULL ELSE $4 END, 
			CASE WHEN $5 = '' THEN NULL ELSE $5 END
		)`
	_, err := s.DB.Exec(query, group, song, releaseDate, text, link)
	return err
}

func (s *SQLSongStore) getSong(group, songName string) (*Song, *uint, error) {
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

func (s *SQLSongStore) deleteSong(id uint) error {
	query := `DELETE tasks WHERE id_song = $1`

	_, err := s.DB.Exec(query, id)
	return err
}

func (s *SQLSongStore) editSong(id int, group, song string, releaseDate *time.Time, text, link *string) error {
	query := `UPDATE tasks SET `
	var args []interface{}
	var updates []string

	if group != "" {
		updates = append(updates, `group = $`+fmt.Sprint(len(args)+1))
		args = append(args, group)
	}
	if song != "" {
		updates = append(updates, `song = $`+fmt.Sprint(len(args)+1))
		args = append(args, song)
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
