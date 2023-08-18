package config

import (
	"os"
)

type config struct {
	DBMySQL       string
	MySQLDBDriver string
	ServerPort    string
}

func GetConfig() *config {
	return &config{
		DBMySQL:       os.Getenv("MYSQL_DB"),
		MySQLDBDriver: os.Getenv("MYSQL_DB_DRIVER"),
		ServerPort:    os.Getenv("SERVER_PORT"),
	}
}
