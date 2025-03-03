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

	currentDir := os.Getenv("PIPEWIRE_DIR")
	if currentDir == "" {
		fmt.Println("No pipewire configuration directory set.")
		return
	} else {
		fmt.Println("Current directory is:", currentDir)
	}
}
