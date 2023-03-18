package models

import (
	"github.com/zemartins81/apiGoPostgres/db"
)

// Update atualiza uma tarefa existente na tabela "todos" com base no ID
func Update(id int64, todo Todo) (int64, error) {
	// Abre uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	// Garante que a conexão será fechada após o término da função
	defer conn.Close()

	// Define a consulta SQL para atualizar uma linha existente na tabela "todos" com base no ID
	sql := `UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4`

	// Executa a consulta SQL com os novos valores da tarefa e o ID correspondente
	res, err := conn.Exec(sql, todo.Title, todo.Description, todo.Done, todo.ID)

	if err != nil {
		return 0, err
	}

	// Retorna o número de linhas afetadas pela atualização (deve ser 1) e possíveis erros
	return res.RowsAffected()
}
