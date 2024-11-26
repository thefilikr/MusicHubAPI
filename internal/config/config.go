package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Env string    `yaml:"env"`
	App ConfigApp `yaml:"app"`
	DB  ConfigDB  `yaml:"db"`
}

type ConfigApp struct {
	PortApp uint `yaml:"port_app"`
}

type ConfigDB struct {
	Host     string `yaml:"host_db"`
	Port     uint   `yaml:"port_db"`
	Username string `yaml:"username_db"`
	Password string `yaml:"password_db"`
	NameDB   string `yaml:"name_db"`
}

// TODO Переписать функцуию под более динамическую проверку
func checkIntegrityConfig(config Config) error {
	var fieldsConfig []string

	switch "" {
	case config.Env:
		fieldsConfig = append(fieldsConfig, "ENV")
	case config.DB.Password:
		fieldsConfig = append(fieldsConfig, "db password")
	case config.DB.Username:
		fieldsConfig = append(fieldsConfig, "db username")
	case config.DB.Host:
		fieldsConfig = append(fieldsConfig, "db host")
	}

	if len(fieldsConfig) != 0 {
		return fmt.Errorf("The field is empty: %s", strings.Join(fieldsConfig, ", "))
	}

	return nil

}

func LoadConfigYAML(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("the file could not be read: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode YAML: %w", err)
	}

	return &config, nil
}

func getUintParameter(parameter string) (*uint, error) {

	uint64Parameter, err := strconv.ParseUint(parameter, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse parameter %s: %v", parameter, err)
	}

	uintParameter := uint(uint64Parameter)

	return &uintParameter, nil
}

func LoadConfigENV() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	var config Config

	// LOAD CONFIG APP
	config.Env = os.Getenv("ENV")

	appPortStr := os.Getenv("APP_PORT")
	if appPortStr == "" {
		return nil, fmt.Errorf("APP_PORT is not set in .env")
	}

	appPort, err := getUintParameter(appPortStr)
	if err != nil {
		return nil, fmt.Errorf("failed parsing APP_PORT: %v", err)
	}
	config.App.PortApp = *appPort

	// LOAD CONFIG DATABASE
	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		return nil, fmt.Errorf("DB_PORT is not set in .env")
	}
	dbPort, err := getUintParameter(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("failed parsing DB_PORT: %v", err)
	}
	config.DB.Port = *dbPort

	config.DB.Host = os.Getenv("DB_HOST")
	config.DB.Username = os.Getenv("DB_USER")
	config.DB.Password = os.Getenv("DB_PASSWORD")
	config.DB.NameDB = os.Getenv("DB_NAME")

	if err := checkIntegrityConfig(config); err != nil {
		return nil, err
	}

	return &config, nil
}
