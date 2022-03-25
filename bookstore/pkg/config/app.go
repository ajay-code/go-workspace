package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	d, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	
	if err != nil {
		log.Fatalln("error", err)
	}
	
	db = d;
}

func GetDB() *gorm.DB {
	return db;
}