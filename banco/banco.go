package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	stringConexao := "root:1234@/dansa?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {

		fmt.Println("Não foi possível abrir o banco")
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		fmt.Println("Não foi possível pingar o banco")
		return nil, erro
	}

	return db, nil
}
