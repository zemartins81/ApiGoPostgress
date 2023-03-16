package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/zemartins81/apiGoPostgres/models"
	"log"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	// Realiza a conversão da string "id" passada na URL para um tipo int
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parser do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Realiza a busca do registro com o id informado
	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao buscar o registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho da resposta HTTP com o tipo "application/json"
	w.Header().Add("Content-Type", "application/json")

	// Codifica o objeto "todo" para JSON e escreve na resposta HTTP
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
