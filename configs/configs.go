package configs

import (
	"github.com/joho/godotenv"
	"os"
)

var Database DatabaseConfig

type DatabaseConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Path     string
}

// InstantiateConfigurationService instantiaton function for config services
func InitializeConf(envPath string) {
	// Load envs
	if err := godotenv.Load(envPath); err != nil {
		panic(err)
	}

	// Database envs
	Database.Path = os.Getenv("DB_PATH")
}
