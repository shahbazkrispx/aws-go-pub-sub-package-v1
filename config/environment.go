package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

/*
*take a key:  type @string
* @return type @string
 */

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
