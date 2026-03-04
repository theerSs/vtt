package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvKey string

func (key EnvKey) GetValue() string {
	return os.Getenv(string(key))
}

const (
	APIPort          EnvKey = "API_PORT"
	PostgresDb       EnvKey = "POSTGRES_DB"
	PostgresUser     EnvKey = "POSTGRES_USER"
	PostgresPassword EnvKey = "POSTGRES_PASSWORD"
	PostgresPort     EnvKey = "POSTGRES_PORT"
)

func Load() error {
	return godotenv.Load(".env")
}
