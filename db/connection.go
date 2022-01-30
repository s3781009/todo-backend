package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func goDotEnvVariable(key string) (string, error) {

	// load .env file
	err := godotenv.Load(".env")

	return os.Getenv(key), err
}
func Connect() *sql.DB {
	var connection, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected to db")
	return connection
}
