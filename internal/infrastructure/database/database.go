package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conectar() *gorm.DB {
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	return db
}
