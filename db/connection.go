
package db

import (
	"database/sql"
	"fmt"

	"github.com/zemartins81/apiGoPostgres/configs"

	_ "github.com/lib/pq" //Importa o driver PostgreSQL
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB() //Recupera as informações de configuração do banco de dados

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database) //Cria uma string de conexão contendo as informações de configuração

	conn, err := sql.Open("postgres", sc) //Abre uma conexão com o banco de dados PostgreSQL
	if err != nil {
		panic(err) //Panic é usado para interromper o programa em caso de erro na abertura da conexão
	}

	err = conn.Ping() //Testa a conexão com o banco de dados
	//Se houver um erro em "conn.Ping()", o erro será retornado na próxima linha

	return conn, err //Retorna a conexão com o banco de dados e um eventual erro
}
