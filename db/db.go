package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConexaoDB() *sql.DB {
	conexao := fmt.Sprintf("%s:%s@tcp(localhost:3306)/db_loja_golang", os.Getenv("root"), os.Getenv("root"))
	//db, err := sql.Open("mysql", conexao)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_loja_golang")
	if err != nil {
		panic(err.Error())
	}
	return db
}
