package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/zehuxx/python-code-api/router"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	router.StartServer()

}
