package constant

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadenv(filenames ...string) {
	if err := godotenv.Load(filenames...); err != nil {
		log.Fatalf("Error getting .env file, %v", err)
	}
}

func Getenv(key string, filenames ...string) string {
	loadenv(filenames...)
	return os.Getenv(key)
}
