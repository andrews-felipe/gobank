package main

import (
	database "gobank/database"
	 "gobank/routes"

)

/**
 *	Bootstrap
 */
func main() {

	// Iniciando conexão com o banco de dados
	
	// Iniciando rotas
	routes.Initialize()

	database.Initialize()
	
}
