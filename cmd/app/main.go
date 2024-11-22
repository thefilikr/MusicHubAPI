package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

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

func main() {

	config, _ := LoadConfig("./../../configs/config.yaml")

	fmt.Println(config)

	log := setupLogger(config.Env)
	log.Info("Setup Logger")

	db := setupDB(config.DB)

	// TODO starting the server
	// TODO запуск сервера

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	db.Close()
}
