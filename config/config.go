package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	AppPort    string `env:"APP_PORT" envDefault:"8080"`
	DBHost     string `env:"DB_HOST" envDefault:"db"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"secret"`
	DBName     string `env:"DB_NAME" envDefault:"subscriptions_db"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"info"`
}

func Load() (*Config, error) {
	c := &Config{}
	err := env.Parse(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
