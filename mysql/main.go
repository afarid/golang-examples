package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	db *sql.DB
	err error
)

func init() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@(127.0.0.1:3306)/test?parseTime=true", dbUser, dbPassword)
	fmt.Println(dsn)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func main() {
  err = db.Ping()
	if err != nil {
		panic(err)
	}
}
