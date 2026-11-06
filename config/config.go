package config

import "os"

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	LogLevel   string
}

func Load() (Config, error) {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	conf := Config{port, "db", "5432", "postres", "secret", "subscriptions_db", "info"}
	return conf, nil
}
