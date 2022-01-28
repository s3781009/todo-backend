package db

import (
	"database/sql"
	"fmt"
	"todo-server/middleware"
	"todo-server/models"
)

func CreateItem(connection *sql.DB, item models.Item) {
	_, err := connection.Query("INSERT INTO todo.todos(id, text, completed,`date-created`,`user-id`) VALUES (?,?,?,?,?)",
		item.Datetime+middleware.Claims["sub"].(string), item.Text, item.Completed, item.Datetime, middleware.Claims["sub"])
	if err != nil {
		fmt.Println("could not create item in db")
		panic(err.Error())
	}
}

func ReadItem(connection *sql.DB) []models.Item {

	rows, _ := connection.Query("SELECT text, completed, `date-created` FROM todo.todos WHERE `user-id`= ?", middleware.Claims["sub"])
	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Text, &item.Completed, &item.Datetime)
		if err != nil {
			fmt.Println("could not read item from db")
			panic(err.Error())
		}
		items = append(items, item)
	}
	return items
}

func UpdateItem(db *sql.DB) {
	//todo
}
func DeleteItem(connection *sql.DB, item models.Item) {
	_, err := connection.Query("DELETE FROM todo.todos WHERE id = ?", item.Datetime+middleware.Claims["sub"].(string))
	if err != nil {
		fmt.Println(err.Error())
	}
}
