package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	stmt.Exec("limubei", "connyun", "2018-04-17")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
