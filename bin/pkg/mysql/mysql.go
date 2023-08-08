package mysql

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vier21/simrs-cdc-monitoring/config"
)

var DB *sqlx.DB

func InitDB() {
	db, err := sqlx.Connect(config.GetConfig().MySQLDBDriver, config.GetConfig().DBMySQL)

	if err != nil {
		log.Fatal(err)
	}
	if db.Ping() != nil {
		log.Fatal(db.Ping())
	}

	DB = db
}

func Close() {
	err := DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
