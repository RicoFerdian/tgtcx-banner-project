package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var (
		err error
	)
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	elephant := "postgres://nrthuxww:vU-_3XInZ2Z5LSIqAwjVUUeGcTJsrLQK@john.db.elephantsql.com/nrthuxww"
	DB, err = sql.Open("postgres", elephant)
	// DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func GetDB() *sql.DB {
	if DB == nil {
		InitDB()
	}
	return DB
}
