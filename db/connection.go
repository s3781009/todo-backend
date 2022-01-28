package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	var connection, err = sql.Open("mysql", "root:2001@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected to db")
	return connection
}
