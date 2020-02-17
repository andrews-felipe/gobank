package domain

import (
	"encoding/json"
	"gobank/models"
	"net/http"

)

/**
 *	Retorna todas as transferencias já realizadas
 */
func getTransfers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transfers []Transfers
	db.Preload("Transfer").Find(&transfers)
	json.NewEncoder(w).Encode(orders)
}
