package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/henrioseptiano/taptalk-diary/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot Connect to Database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("Cannot Connect to Database")
	}
	defer sqlDB.Close()
	app := fiber.New()
	routes.SwaggerRoutes(app)
	routes.UserRoutes(app, db)
	routes.DiaryRoutes(app, db)
	app.Listen(":" + port)
}
