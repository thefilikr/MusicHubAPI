package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"test-task-filikr/internal/service"
)

type SongHandler struct {
	SongService service.SongService
	Log         *slog.Logger
}

// @Description Details about a song
type Song struct {
	Group       string    `json:"group" example:"The Beatles"`
	Song        string    `json:"song" example:"Hey Jude"`
	ReleaseDate *string   `json:"release_date,omitempty" example:"1968-08-26"`
	Text        *[]string `json:"text,omitempty" example:["Na-na-na", "na-na-na-na"]`
	Link        *string   `json:"link,omitempty" example:"https://example.com/song"`
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

func checkRequiredFields(group, song string) error {
	if group == "" || song == "" {
		return errors.New("missing song Group or name song")
	}
	return nil
}

// @Summary Create a new song
// @Description Create a new song with all required fields
// @Tags song
// @Accept json
// @Produce json
// @Param song body Song true "Song object"
// @Success 201 {string} string "Created"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to create song"
// @Router /create [post]
func (h *SongHandler) CreateSongHandler(w http.ResponseWriter, r *http.Request) {

	h.Log.Info("Request create song")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h.Log.Debug("Successful check method")

	song, err := encodeSongJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	h.Log.Debug("Successful encode JSON song")

	if err := checkRequiredFields(song.Group, song.Song); err != nil {
		http.Error(w, "Missing song Group or name song", http.StatusBadRequest)
		return
	}
	h.Log.Debug("Successful check Required Fields")

	if err := h.SongService.CreateSong(song.Group, song.Song, song.ReleaseDate, song.Text, song.Link); err != nil {
		http.Error(w, "Failed to create song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// @Summary Get information about the song.
// @Description Returns information about the song by name and group name.
// @Tags song
// @Accept json
// @Produce json
// @Param group query string true "Name of the group"
// @Param song query string true "The name of the song"
// @Success 200 {object} Song "Song object"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /info [get]
func (h *SongHandler) GetSongHandler(w http.ResponseWriter, r *http.Request) {

	h.Log.Info("Request get song")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h.Log.Debug("Successful check method")

	groupName := r.URL.Query().Get("group")
	songName := r.URL.Query().Get("song")

	if err := checkRequiredFields(groupName, songName); err != nil {
		http.Error(w, "Missing song Group or name song", http.StatusBadRequest)
		return
	}

	h.Log.Debug("Successful check required fields")

	countVerses := r.URL.Query().Get("countVerses")
	numPage := r.URL.Query().Get("numPage")

	song, err := h.SongService.GetSong(groupName, songName, countVerses, numPage)

	if err != nil && song == nil {
		http.Error(w, "Failed get song", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}

// @Summary Delete information about a song.
// @Description Deletes information about the song by name and group name.
// @Tags song
// @Accept json
// @Produce json
// @Param group query string true "Name of the group"
// @Param song query string true "The name of the song"
// @Success 200 {string} string "Song successfully deleted"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /delete [delete]
func (h *SongHandler) DeleteSongHandler(w http.ResponseWriter, r *http.Request) {

	h.Log.Info("Request delete song")

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h.Log.Debug("Successful check method")

	groupName := r.URL.Query().Get("group")
	songName := r.URL.Query().Get("song")

	if err := checkRequiredFields(groupName, songName); err != nil {
		http.Error(w, "Missing song Group or name song", http.StatusBadRequest)
		return
	}

	h.Log.Debug("Successful check required fields")

	if err := h.SongService.DeleteSong(groupName, songName); err != nil {
		http.Error(w, "Failed delete song", http.StatusInternalServerError)

		h.Log.Debug("Failed delete song: ", err)

		return
	}

	h.Log.Debug("Successful delete song")

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Change the information about the song.
// @Description Changes the information about the song by name and group name.
// @Tags song
// @Accept json
// @Produce json
// @Param song body Song true "Song object"
// @Success 200 {string} string "Song successfully changed"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /edit [put]
func (h *SongHandler) EditSongHandler(w http.ResponseWriter, r *http.Request) {

	h.Log.Info("Request edit song")

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h.Log.Debug("Successful check method")

	song, err := encodeSongJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	h.Log.Debug("Successful encode JSON song")

	if err := checkRequiredFields(song.Group, song.Song); err != nil {
		http.Error(w, "Missing song Group or name song", http.StatusBadRequest)
		return
	}

	h.Log.Debug("Successful check required fields")

	if err := h.SongService.EditSong(song.Group, song.Song, song.ReleaseDate, song.Text, song.Link); err != nil {
		h.Log.Debug("Failed edit song: ", err.Error())
		http.Error(w, "Failed edit song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
