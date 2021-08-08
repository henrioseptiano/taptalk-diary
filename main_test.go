package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/henrioseptiano/taptalk-diary/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/gofiber/fiber/v2"
	//utils "github.com/gofiber/v2/utils"
	//user "github.com/henrioseptiano/taptalk-diary/app/user"
)

func Setup() *fiber.App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error : Cannot Loading .env File")
	}

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
	/*sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("Cannot Connect to Database")
	}*/
	//defer sqlDB.Close()
	app := fiber.New()
	routes.SwaggerRoutes(app)
	routes.UserRoutes(app, db)
	routes.DiaryRoutes(app, db)

	return app
}
func testRegister(t *testing.T) {
	app := Setup()
	registerPostBody := map[string]interface{}{
		"birthday": "01-09-1991",
		"email":    "hohoa@hihie.com",
		"fullname": "heheheh",
		"password": "Revian123!",
		"username": "henrio45",
	}
	body, _ := json.Marshal(registerPostBody)
	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewReader(body))
	res, _ := app.Test(req)

	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
	//utils.AssertEqual(t, nil, err, "app.Test")
}

func testLogin(t *testing.T) {
	app := Setup()
	loginPostBody := map[string]interface{}{
		"deviceID": "windows3",
		"password": "revian123",
		"username": "henrio2",
	}
	body, _ := json.Marshal(loginPostBody)
	req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewReader(body))
	res, _ := app.Test(req)

	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}

func testCreateDiaries(t *testing.T) {
	app := Setup()
	loginPostBody := map[string]interface{}{
		"bodyText": "hi",
		"datePost": "08-08-2021",
		"title":    "test2",
	}
	body, _ := json.Marshal(loginPostBody)
	req, _ := http.NewRequest("POST", "/api/v1/diary/create", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjMsInVzZXJuYW1lIjoiaGVucmlvMiJ9.A7dbpb2Bx6d_UZcGZUsNQVj0bG4QF84n29Q-kvrW0TQ")

	res, _ := app.Test(req)
	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}

func testUpdateDiaries(t *testing.T) {
	app := Setup()
	loginPostBody := map[string]interface{}{
		"bodyText": "hi",
		"datePost": "08-08-2021",
		"title":    "test2",
	}
	body, _ := json.Marshal(loginPostBody)
	req, _ := http.NewRequest("PUT", "/api/v1/diary/update/4", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjMsInVzZXJuYW1lIjoiaGVucmlvMiJ9.A7dbpb2Bx6d_UZcGZUsNQVj0bG4QF84n29Q-kvrW0TQ")

	res, _ := app.Test(req)
	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}

func testDeleteDiaries(t *testing.T) {
	app := Setup()
	req, _ := http.NewRequest("DELETE", "/api/v1/diary/delete/4", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjMsInVzZXJuYW1lIjoiaGVucmlvMiJ9.A7dbpb2Bx6d_UZcGZUsNQVj0bG4QF84n29Q-kvrW0TQ")

	res, _ := app.Test(req)
	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}

func testListAllDiaries(t *testing.T) {
	app := Setup()
	req, _ := http.NewRequest("GET", "/api/v1/diary/listall", nil)
	q := req.URL.Query()
	q.Add("year", "2021")
	q.Add("quarter", "3")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjMsInVzZXJuYW1lIjoiaGVucmlvMiJ9.A7dbpb2Bx6d_UZcGZUsNQVj0bG4QF84n29Q-kvrW0TQ")
	res, _ := app.Test(req)
	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}

func testGetDiariesByID(t *testing.T) {
	app := Setup()
	req, _ := http.NewRequest("GET", "/api/v1/diary/3", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjMsInVzZXJuYW1lIjoiaGVucmlvMiJ9.A7dbpb2Bx6d_UZcGZUsNQVj0bG4QF84n29Q-kvrW0TQ")
	res, _ := app.Test(req)
	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}

func testGetCurrentDeviceID(t *testing.T) {
	app := Setup()
	req, _ := http.NewRequest("GET", "/api/v1/getcurrentdeviceid", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjMsInVzZXJuYW1lIjoiaGVucmlvMiJ9.A7dbpb2Bx6d_UZcGZUsNQVj0bG4QF84n29Q-kvrW0TQ")
	res, _ := app.Test(req)
	utils.AssertEqual(t, 200, res.StatusCode, "Status code")
}

func TestRoutes(t *testing.T) {
	testRegister(t)
	testLogin(t)
	testCreateDiaries(t)
	testUpdateDiaries(t)
	testDeleteDiaries(t)
	testListAllDiaries(t)
	testGetDiariesByID(t)
	testGetCurrentDeviceID(t)
}
