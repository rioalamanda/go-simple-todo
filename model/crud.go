package model

import (
	"fmt"

	"github.com/rioalamanda/go-simple-todo/views"
)

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO todo VALUES(?, ?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	insertQ, err := con.Query("DELETE FROM todo WHERE name=?", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM todo WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
