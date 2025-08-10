package enviroment

import (
	"log"

	"github.com/joho/godotenv"
)

// Load loads .env configuration data
func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
