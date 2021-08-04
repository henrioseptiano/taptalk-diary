package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error : Cannot Loading .env File")
	}

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Listen(port)
}
