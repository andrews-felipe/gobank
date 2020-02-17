package models

import (
	"encoding/json"
	"strconv"
	"time"

)

type Transfer struct {
	TransferID         uint      `json:"id" gorm:"primary_key"`
	AccountOrigin      Account   `json:"gorm:"association_foreignkey:AccountID"`
	AccountDestination Account   `json:"gorm:"association_foreignkey:AccountID"`
	Amount             uint      `json:"amount"`
	CreatedAt          time.Time `json:"created_at"`
}
