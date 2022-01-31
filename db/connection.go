package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func Connect() *sql.DB {
	var connection, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected to db")
	return connection
}
