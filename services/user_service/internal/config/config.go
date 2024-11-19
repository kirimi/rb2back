package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var cfg *Config

type Config struct {
	DB   DB   `yaml:"db"`
	Grpc Grpc `yaml:"grpc"`
}

func GetConfig(yamlFilePath string) Config {
	if cfg != nil {
		return *cfg
	}

	cfg := &Config{
		DB:   NewDBWithDefaults(),
		Grpc: NewGrpcWithDefaults(),
	}

	file, err := os.Open(filepath.Clean(yamlFilePath))
	if err != nil {
		log.Warn().Err(err)
		return *cfg
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Warn().Err(err)
		return *cfg
	}

	return *cfg
}
