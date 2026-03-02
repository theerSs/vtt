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
 APIPort EnvKey = "API_PORT" 
)

func Load() error {
	return godotenv.Load(".env")
}