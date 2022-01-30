package db

import (
	"database/sql"
	"fmt"
	"todo-server/models"
)

func CreateItem(connection *sql.DB, item models.Item, token string) {
	_, err := connection.Query("INSERT INTO todo.todos(id, text, completed,`date-created`,`user-id`) VALUES (?,?,?,?,?)",
		item.Datetime+token, item.Text, item.Completed, item.Datetime, token)
	if err != nil {
		fmt.Println("could not create item in db")
		panic(err.Error())
	}
}

func ReadItem(connection *sql.DB, token string) []models.Item {

	rows, _ := connection.Query("SELECT text, completed, `date-created` FROM todo.todos WHERE `user-id`= ?", token)
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

func UpdateItem(connection *sql.DB, item models.Item, token string) {
	//todo
	_, err := connection.Query("UPDATE todo.todos SET `text` = ? WHERE id = ?", item.Text, item.Datetime+token)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func DeleteItem(connection *sql.DB, item models.Item, token string) {
	_, err := connection.Query("DELETE FROM todo.todos WHERE id = ?", item.Datetime+token)
	if err != nil {
		fmt.Println(err.Error())
	}
}
