package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var MONGO_URI = os.Getenv("MONGO_URI")