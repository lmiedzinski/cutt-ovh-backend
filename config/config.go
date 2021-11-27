package config

import "os"

type Config struct {
	PortNumber               string
	BaseUrl                  string
	LogLevel                 string
	PostgresConnectionString string
}

func GetConfig() *Config {
	return &Config{
		PortNumber:               getEnv("CFG_PORTNUMBER", "9000"),
		BaseUrl:                  getEnv("CFG_BASEURL", "/"),
		LogLevel:                 getEnv("CFG_LOGLEVEL", "info"),
		PostgresConnectionString: getEnv("CFG_POSTGRESCONNECTIONSTRING", "postgres://postgres:mysecretpassword@localhost:5432/cutt_ovh_database"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
