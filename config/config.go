package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Load key (if found) from the .env file.
func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load .env file: %v\n", err)
	}

	return os.Getenv(key)
}
