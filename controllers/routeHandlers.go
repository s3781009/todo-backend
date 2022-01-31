package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-server/db"
	"todo-server/models"
)

func GetItems(connection *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Get("token")

		items := db.ReadItem(connection, token.(string))

		c.JSON(200, items)

	}
}
func PostItems(connection *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.Item
		err := c.Bind(&item)
		if err != nil {
			c.JSON(500, "could not bind request with model")
		}
		c.JSON(200, item)
		token, _ := c.Get("token")
		db.CreateItem(connection, item, token.(string))
	}
}

func UpdateItem(connection *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.Item
		token, _ := c.Get("token")
		err := c.Bind(&item)
		if err != nil {
			fmt.Println(err.Error())

		}
		db.UpdateItem(connection, item, token.(string))
	}
}

func DeleteItem(connection *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.Item
		err := c.Bind(&item)
		if err != nil {
			c.JSON(500, err.Error())
		}
		token, _ := c.Get("token")
		db.DeleteItem(connection, item, token.(string))
	}
}
