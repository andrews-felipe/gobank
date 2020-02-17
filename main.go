package main

import (
	"gobank/database"
	"gobank/models"
	"gobank/routes"
	"log"
	"net/http"

)

/**
 *	Bootstrap
 */
func main() {

	// Iniciando rotas
	routes.initialize()

	// Iniciando conexão com o banco de dados
	database.initialize()

	log.Fatal(http.ListenAndServe(":8080", router))
}
