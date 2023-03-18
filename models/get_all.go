package models

import "github.com/zemartins81/apiGoPostgres/db"

// GetAll busca todas as tarefas da tabela "todos"
func GetAll() (todos []Todo, err error) {

	// Abre uma conexão com o banco de dados
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	// Garante que a conexão será fechada após o término da função
	defer conn.Close()

	// Executa uma consulta SQL para obter todas as linhas da tabela "todos"
	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	// Itera sobre todas as linhas retornadas pela consulta
	for rows.Next() {
		var todo Todo

		// Preenche a estrutura "todo" com os dados da linha atual
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue // Se ocorrer um erro, ignora a linha e continua a iteração
		}

		// Adiciona a tarefa à lista de tarefas encontradas
		todos = append(todos, todo)
	}

	// Retorna a lista de tarefas encontradas (ou uma lista vazia) e possíveis erros
	return
}

