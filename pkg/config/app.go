package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// here we connect with db
var db *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=Cicada_3301 dbname=book_management port=5000 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error", err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
