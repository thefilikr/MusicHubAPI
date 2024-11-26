package app

import (
	"database/sql"
	"log/slog"
	"net/http"
	"strconv"
	"test-task-filikr/internal/server/http/handler"
	"test-task-filikr/internal/service"
	"test-task-filikr/internal/store"
)

func SetupApp(db *sql.DB, log *slog.Logger, port uint) {

	songStore := &store.SQLSongStore{
		DB:  db,
		Log: log,
	}

	log.Debug("Init layer store")

	songService := service.SongService{
		SongStore: songStore,
		Log:       log,
	}

	log.Debug("Init layer service")

	songHandler := handler.SongHandler{
		SongService: songService,
		Log:         log,
	}

	log.Debug("Init layer handler")

	http.HandleFunc("/songs", songHandler.GetSongsHandler)
	http.HandleFunc("/song/info", songHandler.GetSongHandler)
	http.HandleFunc("/song/create", songHandler.CreateSongHandler)
	http.HandleFunc("/song/delete", songHandler.DeleteSongHandler)
	http.HandleFunc("/song/edit", songHandler.EditSongHandler)

	log.Debug("Init handlers")

	if err := http.ListenAndServe(":"+strconv.Itoa(int(port)), nil); err != nil {
		log.Error("Server failed to start: ", err)
	}
}
