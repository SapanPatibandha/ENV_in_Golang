package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("HOME: ", os.Getenv("HOME"))

	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		fmt.Println("the env var SHELL is not set")
	} else {
		fmt.Println("SHELL: ", shell)
	}

	err := os.Setenv("NAME", "Sapan Patibandha")
	if err != nil {
		fmt.Println("Could not set the env var NAME")
	}

	fmt.Printf("Name: %s\n", os.Getenv("NAME"))

	envErr := godotenv.Load(".env")

	if envErr != nil {
		fmt.Println("Could not load env file")
		os.Exit(1)
	}

	fmt.Printf("API_KEY %s\n", os.Getenv("API_KEY"))
	fmt.Printf("DB_PASS %s\n", os.Getenv("DB_PASS"))

	envMap, mapErr := godotenv.Read(".env")

	if mapErr != nil {
		fmt.Printf("Error loading .env into map[string]string\n")
		os.Exit(1)
	}

	fmt.Printf("API_KEY from godotenv %s\n", envMap["API_KEY"])
	fmt.Printf("DB_PASS from godotenv %s\n", envMap["DB_PASS"])
}
