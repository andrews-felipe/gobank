package database

import (
	"fmt"
	"gobank/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var db *gorm.DB

/**
 *	Inicialização do banco de dados
 */
func Initialize() {
	// Iniciando banco de dados
	var err error
	dataSourceName := "root:123456@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("falha na conexão com o banco de dados")
	}

	// Criando bando de dados
	db.Exec("CREATE DATABASE gobank_db")
	db.Exec("USE gobank_db")

	// migration pra criação dos schemas
	db.AutoMigrate(&models.Account{}, &models.Transfer{})
}
