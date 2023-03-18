package models

import "github.com/zemartins81/apiGoPostgres/db"

// Get busca uma tarefa com base em seu ID
func Get(id int64) (todo Todo, err error) {

	// Abre uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	// Garante que a conexão será fechada após o término da função
	defer conn.Close()

	// Executa uma consulta SQL para obter uma linha da tabela "todos"
	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	// Preenche a estrutura "todo" com os dados retornados pela consulta
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	// Retorna a tarefa encontrada (ou uma tarefa vazia) e possíveis erros
	return
}

