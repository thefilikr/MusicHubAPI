package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"test-task-filikr/internal/service"
	"time"
)

type SongHandler struct {
	SongService service.SongService
}

type Song struct {
	Group       string     `json:"group"`
	Song        string     `json:"song"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Text        *string    `json:"text,omitempty"`
	Link        *string    `json:"link,omitempty"`
}

// // in a good way, such functionality should be transferred to another file, since json decoding will occur in other handlers as well
// по хорощему такой функционал надо выносить в другой файл, т.к. декдирование json будет происходить и в дргуих хендлерах
func encodeSongJSON(songJSON io.ReadCloser) (*Song, error) {
	var song Song
	if err := json.NewDecoder(songJSON).Decode(&song); err != nil {
		return nil, err
	}

	return &song, nil
}

func checkRequiredFields(song *Song) error {
	if song.Group == "" || song.Song == "" {
		return errors.New("missing song Group or name song")
	}
	return nil
}

func (h *SongHandler) CreateSongHandler(w http.ResponseWriter, r *http.Request) {

	song, err := encodeSongJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := checkRequiredFields(song); err != nil {
		http.Error(w, "Missing song Group or name song", http.StatusBadRequest)
		return
	}

	if err := h.SongService.CreateSong(song.Group, song.Song, song.ReleaseDate, song.Text, song.Link); err != nil {
		http.Error(w, "Failed to create song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SongHandler) GetSong(w http.ResponseWriter, r *http.Request) {
	nameSong, err := encodeSongJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := checkRequiredFields(nameSong); err != nil {
		http.Error(w, "Missing song Group or name song", http.StatusBadRequest)
		return
	}

	song, _, err := h.SongService.GetSong(nameSong.Group, nameSong.Song)

	if err != nil {
		http.Error(w, "Failed get song", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

func (h *SongHandler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	song, err := encodeSongJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := checkRequiredFields(song); err != nil {
		http.Error(w, "Missing song Group or name song", http.StatusBadRequest)
		return
	}

	if err := h.SongService.DeleteSong(song.Group, song.Song); err != nil {
		http.Error(w, "Failed delete song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SongHandler) EditSong(w http.ResponseWriter, r *http.Request) {
	song, err := encodeSongJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := checkRequiredFields(song); err != nil {
		http.Error(w, "Missing song Group or name song", http.StatusBadRequest)
		return
	}

	if err := h.SongService.EditSong(song.Group, song.Song, song.ReleaseDate, song.Text, song.Link); err != nil {
		http.Error(w, "Failed edit song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
