package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zemartins81/apiGoPostgres/models"
)

// Handler que recebe uma requisição para criar uma nova tarefa
func Create(w http.ResponseWriter, r *http.Request) {

	// Decodifica o corpo da requisição em um objeto do tipo models.Todo
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		// Se ocorrer algum erro na decodificação, retorna um erro HTTP 500 com uma mensagem de erro genérica
		log.Printf("Erro ao fazer o decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Insere a nova tarefa no banco de dados
	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		// Se ocorrer algum erro na inserção, retorna uma mensagem de erro contendo a descrição do erro
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Erro ao tentar inserir: %v", err),
		}
	} else {
		// Se a inserção for bem-sucedida, retorna uma mensagem de sucesso contendo o ID da nova tarefa
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso! ID: %d", id),
		}
	}

	// Monta uma resposta HTTP em JSON com o resultado da operação (sucesso ou erro)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

