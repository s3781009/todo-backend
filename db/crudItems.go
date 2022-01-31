package db

import (
	"database/sql"
	"fmt"
	"log"
	"todo-server/models"
)

func CreateItem(connection *sql.DB, item models.Item, token string) {
	e := connection.Ping()
	if e != nil {
		fmt.Println("could not connect to db")
	}
	_, err := connection.Query("INSERT INTO todos.todolist(id, text, completed, date_created, user_id) VALUES ($1,$2,$3,$4,$5)",
		item.Datetime+token, item.Text, item.Completed, item.Datetime, token)
	if err != nil {
		log.Fatal("cannot write to db ", err)

	}
}

func ReadItem(connection *sql.DB, token string) []models.Item {

	rows, _ := connection.Query("SELECT text, completed, date_created FROM todos.todolist WHERE user_id= $1", token)
	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Text, &item.Completed, &item.Datetime)
		if err != nil {
			log.Fatal("cannot read from db", err)
		}
		items = append(items, item)
	}
	return items
}

func UpdateItem(connection *sql.DB, item models.Item, token string) {
	//todo
	_, err := connection.Query("UPDATE todos.todolist SET text = $1 WHERE id = $2", item.Text, item.Datetime+token)
	if err != nil {
		log.Fatal("cannot update db", err)
	}
}
func DeleteItem(connection *sql.DB, item models.Item, token string) {
	_, err := connection.Query("DELETE FROM todos.todolist WHERE id = $1", item.Datetime+token)
	if err != nil {
		log.Fatal("cannot delete from db", err)
	}
}
