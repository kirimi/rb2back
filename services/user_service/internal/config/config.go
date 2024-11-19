package config

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"time"
)

var cfg *Config

type DB struct {
	DSN             string        `yaml:"DSN"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"`
	ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
}

type Config struct {
	DB DB `yaml:"db"`
}

func (d *DB) GetDSN() string {
	return d.DSN
}

func (d *DB) GetMaxOpenConns() int {
	return d.MaxOpenConns
}

func (d *DB) GetMaxIdleConns() int {
	return d.MaxIdleConns
}

func (d *DB) GetConnMaxIdleTime() time.Duration {
	return d.ConnMaxIdleTime
}

func (d *DB) GetConnMaxLifeTime() time.Duration {
	return d.ConnMaxLifeTime
}

func GetConfig(yamlFilePath string) Config {
	if cfg != nil {
		return *cfg
	}

	cfg := &Config{
		DB: DB{
			DSN:             "postgres://postgres:postrges@locahost:5432/db",
			MaxOpenConns:    5,
			MaxIdleConns:    5,
			ConnMaxIdleTime: time.Minute * 5,
			ConnMaxLifeTime: time.Minute * 5,
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
