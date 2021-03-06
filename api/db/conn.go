package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:root@tcp(localhost:3307)/uraban?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
