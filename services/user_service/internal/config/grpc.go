package config

import "time"

type Grpc struct {
	Host              string        `yaml:"host"`
	Port              int           `yaml:"port"`
	MaxConnectionIdle time.Duration `yaml:"maxConnectionIdle"`
	MaxConnectionAge  time.Duration `yaml:"maxConnectionAge"`
	Timeout           time.Duration `yaml:"timeout"`
}

func NewGrpcWithDefaults() Grpc {
	return Grpc{
		Host:              "localhost",
		Port:              8082,
		MaxConnectionIdle: time.Minute * 5,
		MaxConnectionAge:  time.Minute * 5,
		Timeout:           time.Second * 15,
	}
}

func (g *Grpc) GetHost() string {
	return g.Host
}

func (g *Grpc) GetPort() int {
	return g.Port
}

func (g *Grpc) GetMaxConnectionIdle() time.Duration {
	return g.MaxConnectionIdle
}

func (g *Grpc) GetMaxConnectionAge() time.Duration {
	return g.MaxConnectionAge
}

func (g *Grpc) GetTimeout() time.Duration {
	return g.Timeout
}
