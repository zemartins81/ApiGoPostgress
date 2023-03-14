package models

import "github.com/zemartins81/apiGoPostgres/db"

func GetAll() (todos []Todo, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.QueryRow(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue //ideal seria um log
		}

		todos = append(todos, todo)
	}

	return
}