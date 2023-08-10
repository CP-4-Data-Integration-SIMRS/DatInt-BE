package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitMysqlDB() (err error) {
	dsn := "root@tcp(127.0.0.1:3306)/"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	return
}

func Close() {
	err := DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
