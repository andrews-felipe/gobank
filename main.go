package main

import (
	"fmt"
	"gobank/database"
	"gobank/models"
	"gobank/routes"
	"net/http"

)

var db *gorm.DB

/**
 *	Inicialização do banco de dados
 */
func main() {

	// Iniciando rotas
	routes.initialize()

	// Iniciando conexão com o banco de dados
	database.initialize()

	//
	log.Fatal(http.ListenAndServe(":8080", router))
}
