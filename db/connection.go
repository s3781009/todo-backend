package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func dbEnv(key string) string {

	env, _ := godotenv.Read(".env")
	secret := env[key]
	return secret
}
func Connect() *sql.DB {

	var connection, err = sql.Open("postgres", dbEnv("DATABASE_URL"))
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	err = connection.Ping()
	if err != nil {
		log.Fatal("cannot ping", err)
	}
	
	return connection
}
