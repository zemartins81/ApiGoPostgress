package models

import (
	"github.com/zemartins81/apiGoPostgres/db"
)

// Insert insere uma nova tarefa na tabela "todos"
func Insert(todo Todo) (id int64, err error) {

	// Abre uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	// Garante que a conexão será fechada após o término da função
	defer conn.Close()

	// Define a consulta SQL para inserir uma nova linha na tabela "todos" e obter o ID da nova linha inserida
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	// Executa a consulta SQL e armazena o ID da nova linha inserida na variável "id"
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	// Retorna o ID da nova linha inserida (ou zero) e possíveis erros
	return
}

