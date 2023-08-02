package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

	db, err := sqlx.Connect("mysql", "root@tcp(127.0.0.1:3306)/" + "coba")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	p := []string{}

	db.Select(&p, "SHOW TABLES")
	for i := range p {

		log.Println(p[i])
	}
}


