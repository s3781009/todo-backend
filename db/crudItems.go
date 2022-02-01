package db

import (
	"database/sql"
	"fmt"
	"log"
	"todo-server/models"
)

func CreateItem(connection *sql.DB, item models.Item, token string) {
	_, err := connection.Query("INSERT INTO todos.todolist(id, text, completed,date_created,user_id) VALUES ($1,$2,$3,$4,$5)",
		item.Datetime+token, item.Text, item.Completed, item.Datetime, token)
	if err != nil {
		log.Fatal("could not create item in db", err)
	}
}

func ReadItem(connection *sql.DB, token string) []models.Item {

	rows, e := connection.Query("SELECT text, completed, date_created FROM todos.todolist  WHERE user_id= $1", token)
	if e!=nil{
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

func UpdateItem(connection *sql.DB, item models.Item, token string) {
	//todo
	_, err := connection.Query("UPDATE todos.todolist SET text = $1 WHERE id = $2", item.Text, item.Datetime+token)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func DeleteItem(connection *sql.DB, item models.Item, token string) {
	_, err := connection.Query("DELETE FROM  todos.todolist WHERE id = $1", item.Datetime+token)
	if err != nil {
		log.Fatal("could not delete item from db", err)
	}
}
