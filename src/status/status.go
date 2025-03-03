package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	currentBuffer := os.Getenv("PREV_BUFFER")
	if currentBuffer == "" {
		fmt.Println("No previous buffer size set.")
		return
	} else {
		fmt.Println("Current buffer size is:", currentBuffer)
	}
}
