package db

import (
	"database/sql"
	"log"
	"todo-server/models"
)

// CreateItem create a to do item in the databse which  is used by the POST route handler
func CreateItem(connection *sql.DB, item models.Item, token string) {
	_, err := connection.Query("INSERT INTO todos.todolist(id, text, completed,date_created,user_id) VALUES ($1,$2,$3,$4,$5)",
		item.Datetime+token, item.Text, item.Completed, item.Datetime, token)
	if err != nil {
		log.Fatal("could not create item in db", err)
	}
}

// ReadItem read  the list of to do items for a user  which is passed to the correlating GET handler
func ReadItem(connection *sql.DB, token string) []models.Item {

	rows, e := connection.Query("SELECT text, completed, date_created FROM todos.todolist  WHERE user_id= $1", token)
	if e != nil {
		log.Fatal(e)
	}
	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Text, &item.Completed, &item.Datetime)
		if err != nil {
			log.Fatal("could not read from db", err)
		}
		items = append(items, item)
	}
	return items
}

// UpdateItem Update to do item is the user has either completed or updated the text of the to do item
//and passed to the PATCH route hanlder
func UpdateItem(connection *sql.DB, item models.Item, token string) {
	_, err := connection.Query("UPDATE todos.todolist SET text = $1 ,completed = $2 WHERE id = $3", item.Text, item.Completed, item.Datetime+token)
	if err != nil {
		log.Fatal("cannot update item in db", err)
	}
}

// DeleteItem delete and to do item and used by the DELETE route handler
func DeleteItem(connection *sql.DB, item models.Item, token string) {
	_, err := connection.Query("DELETE FROM  todos.todolist WHERE id = $1", item.Datetime+token)
	if err != nil {
		log.Fatal("could not delete item from db", err)
	}
}
