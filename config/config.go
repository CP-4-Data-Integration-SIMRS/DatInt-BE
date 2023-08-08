package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	DBMySQL       string
	MySQLDBDriver string
	ServerPort    string
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env")
	}
}

func GetConfig() *config {
	return &config{
		DBMySQL:       os.Getenv("MYSQL_DB"),
		MySQLDBDriver: os.Getenv("MYSQL_DB_DRIVER"),
		ServerPort:    os.Getenv("SERVER_PORT"),
	}
}
