package main

import (
	"fmt"
	"net/http"

	"github.com/zemartins81/apiGoPostgres/configs"
	"github.com/zemartins81/apiGoPostgres/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	// Cria um novo roteador HTTP utilizando o pacote "chi"
	r := chi.NewRouter()

	// Registra as rotas para manipular as tarefas "Todo"
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	// Inicia o servidor HTTP na porta especificada nas configurações
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
