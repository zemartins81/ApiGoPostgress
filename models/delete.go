// Pacote models é responsável por fornecer as funções que modelam o acesso e manipulação dos dados da aplicação.

package models

// Importa o pacote "db" que contém a função "OpenConnection" que permite abrir uma conexão com o banco de dados.
import "github.com/zemartins81/apiGoPostgres/db"

// A função Delete é responsável por excluir uma linha da tabela "todos" do banco de dados PostgreSQL.
// Recebe um parâmetro "id" do tipo int64 que identifica a linha a ser excluída.
// Retorna o número de linhas afetadas pela operação e um possível erro caso a operação falhe.
func Delete(id int64) (int64, error) {
	// Abre uma conexão com o banco de dados usando a função OpenConnection do pacote db.
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	// Adia o fechamento da conexão para garantir que a conexão seja sempre fechada, mesmo em caso de erro.
	defer conn.Close()

	// Executa uma operação de exclusão na tabela "todos" do banco de dados, utilizando o identificador passado como parâmetro.
	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	// Retorna o número de linhas afetadas pela operação de exclusão.
	return res.RowsAffected()
}
