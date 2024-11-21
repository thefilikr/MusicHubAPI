package main

import (
	"fmt"
	"log/slog"
	"os"

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
	PortDB   uint   `yaml:"port_db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
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

func main() {

	config, _ := LoadConfig("./../../configs/config.yaml")

	fmt.Println(config)

	// TODO логер

	log := setupLogger(config.Env)
	log.Info("Setup Logger")

	// TODO connect db

	// TODO start app
}
