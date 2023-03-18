
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/zemartins81/apiGoPostgres/models"
)

// Handler que recebe uma requisição para deletar uma tarefa pelo ID
func Delete(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da URL da requisição e converte para inteiro
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parser do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Deleta a tarefa correspondente no banco de dados
	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover o registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Se mais de uma linha foi deletada, um erro é logado
	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}

	// Cria uma resposta HTTP em JSON com uma mensagem de sucesso
	resp := map[string]any{"Error": false, "Message": "Dados removidos com sucesso"}

	// Define o header da resposta como JSON e encoda a resposta em JSON para ser enviada
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
