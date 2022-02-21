package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func environmentVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("%v: Could not load .env file", err)
	}
	return os.Getenv(key)
}

func main() {
	password := environmentVariable("PASSWORD")
	//  "driver" "user:password@/database"
	db, err := sql.Open("mysql", "root:"+password+"/")
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
