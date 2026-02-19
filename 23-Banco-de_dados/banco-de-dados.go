package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Driver import implicito
)

func main() {
	stringConexao := "myuser:secret@/meli_social?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão aberta ")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(linhas)

	defer linhas.Close()

}
