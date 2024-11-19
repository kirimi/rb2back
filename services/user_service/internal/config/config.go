package config

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var cfg *Config

type Database struct {
	Host       string `yaml:"host"`
	Port       int64  `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Migrations string `yaml:"migrations"`
	DBName     string `yaml:"dbname"`
	SslMode    bool   `yaml:"sslmode"`
	Driver     string `yaml:"driver"`
}

type Config struct {
	Database Database `yaml:"database"`
}

func GetConfig(yamlFilePath string) Config {
	if cfg != nil {
		return *cfg
	}

	cfg := &Config{
		Database: Database{
			Host:       "localhost",
			Port:       5432,
			User:       "user",
			Password:   "password",
			DBName:     "postrgres",
			SslMode:    false,
			Migrations: "",
			Driver:     "pgx",
		},
	}

	file, err := os.Open(filepath.Clean(yamlFilePath))
	if err != nil {
		log.Warn(err)
		return *cfg
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Warn(err)
		return *cfg
	}

	return *cfg
}
