package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"todo-server/db"
	"todo-server/models"
)

func GetItems(connection *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		items := db.ReadItem(connection)

		c.JSON(200, items)

	}
}
func PostItems(connection *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.Item
		err := c.Bind(&item)
		if err != nil {
			c.JSON(500, err.Error())
		}
		c.JSON(200, item)
		db.CreateItem(connection, item)
	}
}

func DeleteItem(connection *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.Item
		err := c.Bind(&item)
		if err != nil {
			c.JSON(500, err.Error())
		}
		db.DeleteItem(connection, item)
	}
}
