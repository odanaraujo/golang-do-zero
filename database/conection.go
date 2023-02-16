package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Abre conexão com o banco de dados
func Conection() (*sql.DB, error) {
	stringConnection := "root:1234@/bdestudos?charset=utf8&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", stringConnection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
