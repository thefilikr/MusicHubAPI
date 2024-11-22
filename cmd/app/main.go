package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"database/sql"

	_ "github.com/lib/pq"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env            string    `yaml:"env"`
	App            ConfigApp `yaml:"app"`
	DB             ConfigDB  `yaml:"db"`
	MigrationsPath string    `yaml:"migrations_path"`
}

type ConfigApp struct {
	PortApp uint `yaml:"port_app"`
}

type ConfigDB struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port_db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	NameDB   string `yaml:"name_db"`
}

func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("не удалось декодировать YAML: %w", err)
	}

	return &config, nil

}

const (
	envLocal = "local"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupDB(config ConfigDB) *sql.DB {

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.NameDB,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}

type Song struct {
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Text        []string  `json:"text,omitempty"`
	Link        string    `json:"link,omitempty"`
}

func GetSong(w http.ResponseWriter, r *http.Request) {
}

func CreateSong(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var newSong Song
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newSong); err != nil {
		http.Error(w, "Incorrect JSON", http.StatusBadRequest)
		return
	}

	if newSong.Group == "" || newSong.Song == "" {
		http.Error(w, "Required fields are missing: Group or Song", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]string{"message": "The song was created successfully"}
	json.NewEncoder(w).Encode(response)
}

func EditSong(w http.ResponseWriter, r *http.Request) {
}

func DeleteSong(w http.ResponseWriter, r *http.Request) {
}

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/user", GetSong)
	router.HandleFunc("/users/create", CreateSong)
	router.HandleFunc("/users/edit", EditSong)
	router.HandleFunc("/users/delete", DeleteSong)

	return router
}

func startApp(config ConfigApp) {
	router := NewRouter()

	// log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":"+string(config.PortApp), router); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func main() {

	config, _ := LoadConfig("./../../configs/config.yaml")

	fmt.Println(config)

	log := setupLogger(config.Env)
	log.Info("Setup Logger")

	db := setupDB(config.DB)

	startApp(config.App)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	db.Close()
}
