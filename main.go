package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/henrioseptiano/taptalk-diary/routes"
	"github.com/joho/godotenv"

	_ "github.com/henrioseptiano/taptalk-diary/docs"
)

// @title Fiber TapTalk Diary API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @schemes http
// @securityDefinitions.apikey Token
// @in header
// @name Authorization
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error : Cannot Loading .env File")
	}

	port := os.Getenv("PORT")

	app := fiber.New()
	routes.SwaggerRoutes(app)
	routes.UserRoutes(app)
	app.Listen(":" + port)
}
