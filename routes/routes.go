package routes

import (
	"log"

	"github.com/gorilla/mux"

)

func initialize() {
	router := mux.NewRouter()
	// obtém a lista de transferencias
	router.HandleFunc("/transfers", getTransfers).Methods("GET")
	// faz transferencia de um Account para outro
	router.HandleFunc("/transfers", sendTransfer).Methods("POST")
	// obtém a lista de contas
	router.HandleFunc("/accounts", getAccounts).Methods("GET")
	// obtém o saldo da conta
	router.HandleFunc("/accounts/{account_id}/balance", getBallance).Methods("GET")
	// criar uma conta
	router.HandleFunc("/accounts", createAccount).Methods("POST")

}
