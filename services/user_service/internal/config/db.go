package config

import "time"

type DB struct {
	DSN             string        `yaml:"DSN"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"`
	ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
}

func NewDBWithDefaults() DB {
	return DB{
		DSN:             "postgres://postgres:postrges@locahost:5432/db",
		MaxOpenConns:    5,
		MaxIdleConns:    5,
		ConnMaxIdleTime: time.Minute * 5,
		ConnMaxLifeTime: time.Minute * 5,
	}
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
