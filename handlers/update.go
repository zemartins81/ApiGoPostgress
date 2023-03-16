package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"

	"github.com/zemartins81/apiGoPostgres/models"
)
// Essa é uma função HTTP que recebe uma requisição para atualizar uma tarefa "Todo" existente.
// Ela começa convertendo o parâmetro "id" da URL da requisição para um inteiro utilizando a função Atoi do pacote "strconv".
// Se ocorrer um erro durante a conversão, um erro HTTP 500 é retornado com uma mensagem de erro.
// Em seguida, decodifica o JSON enviado na requisição para um objeto do tipo Todo, utilizando um decoder do pacote "encoding/json".
// Chama o método "Update" do modelo "Todo" para atualizar o objeto no banco de dados.
// Se ocorrer um erro durante o processo de decodificação ou atualização, um erro HTTP 500 é retornado com uma mensagem de erro.
// Se mais de um registro for atualizado durante a operação de atualização, um log de erro é gerado.
// Se tudo ocorrer corretamente, um objeto JSON contendo um indicador de sucesso e uma mensagem de sucesso é retornado.

func Update(w http.ResponseWriter, r *http.Request) {
// Converte o parâmetro "id" da URL da requisição para um inteiro
id, err := strconv.Atoi(chi.URLParam(r, "id"))
if err != nil {
log.Printf("Erro ao fazer parser do id: %v", err)
http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
return
}

// Decodifica o JSON enviado na requisição para um objeto do tipo Todo
var todo models.Todo
err = json.NewDecoder(r.Body).Decode(&todo)
if err != nil {
	log.Printf("Erro ao fazer o decode do json: %v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}

// Chama o método "Update" do modelo "Todo" para atualizar o objeto no banco de dados
rows, err := models.Update(int64(id), todo)
if err != nil {
	log.Printf("Erro ao atualizar o registro: %v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}

// Se mais de um registro for atualizado durante a operação de atualização, um log de erro é gerado
if rows > 1 {
	log.Printf("Error: foram atualizados %d registros", rows)
}

// Cria um objeto de resposta contendo um indicador de sucesso e uma mensagem
resp := map[string]interface{}{
	"Error":   false,
	"Message": "Dados atualizados com sucesso",
}

// Define o tipo de conteúdo da resposta como JSON e codifica o objeto de resposta para JSON
w.Header().Add("Content-Type", "application/json")
json.NewEncoder(w).Encode(resp)
