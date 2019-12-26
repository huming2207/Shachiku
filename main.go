package shachiku

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load configuration: %w", err)
	}
}
