package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	//  "driver" "user:password@/database"
	db, err := sql.Open("mysql", "root:@/")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	exec(db, "create database if not exists dieta")
	exec(db, "use dieta")
	exec(db, "drop table if exists comida")
	exec(db, `create table comida(
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
