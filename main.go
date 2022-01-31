package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"todo-server/db"
	"todo-server/routes"
)

func main() {

	connection := db.Connect()
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.Setup(app, connection)

	err := app.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal("cannot start server", err)
	}
	fmt.Print("finished")
}
