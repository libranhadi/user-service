package database

import (
	"database/sql"
	"service-user/config"
	"time"

	_ "github.com/lib/pq"
)

func InitDatabase() *sql.DB {
	var db *sql.DB
	var err error

	config := config.ConfigDatabase()

	db, err = sql.Open("postgres", config)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
