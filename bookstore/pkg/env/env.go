package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_URI string
	DB_USER string
	DB_PASSWORD string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	DB_URI = os.Getenv("DB_URI")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
}