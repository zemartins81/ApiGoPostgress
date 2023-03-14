package db

import (
	"database/sql"
	"fmt"

	"github.com/zemartins81/apiGoPostgres/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err) //Nunca é uma boa ideia usar isso em produção
	}

	err = conn.Ping()

	return conn, err
}
