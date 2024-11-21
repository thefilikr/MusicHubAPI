package main

import (
	"fmt"
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

func main() {

	config, _ := LoadConfig("./../../configs/config.yaml")

	fmt.Println(config)

	// TODO логер

	// TODO connect db

	// TODO start app
}
