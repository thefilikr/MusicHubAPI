package service

import (
	"fmt"
	"test-task-filikr/internal/store"
	"time"
)

type SongService struct {
	SongStore store.SongStore
}

func (s *SongService) CreateSong(group, song string, releaseDate *time.Time, text, link *string) error {
	return s.SongStore.CreateSong(group, song, releaseDate, text, link)
}

func (s *SongService) GetSong(group, song string) (*store.Song, *uint, error) {
	return s.SongStore.GetSong(group, song)
}

func (s *SongService) DeleteSong(group, song string) error {
	_, id, err := s.SongStore.GetSong(group, song)
	if err != nil {
		return fmt.Errorf("The song was not found")
		// TODO изменить реализацию логирования
	}

	return s.SongStore.DeleteSong(*id)
}

func (s *SongService) EditSong(group, song string, releaseDate *time.Time, text, link *string) error {
	_, id, err := s.SongStore.GetSong(group, song)
	if err != nil {
		return fmt.Errorf("The song was not found")
		// TODO изменить реализацию логирования
	}

	return s.SongStore.EditSong(*id, group, song, releaseDate, text, link)

}
