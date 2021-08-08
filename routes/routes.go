package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	Diary "github.com/henrioseptiano/taptalk-diary/app/diaries/controller"
	DiariesRepo "github.com/henrioseptiano/taptalk-diary/app/diaries/repository"
	DiaryService "github.com/henrioseptiano/taptalk-diary/app/diaries/services"
	User "github.com/henrioseptiano/taptalk-diary/app/users/controller"
	UserRepo "github.com/henrioseptiano/taptalk-diary/app/users/repository"
	UserService "github.com/henrioseptiano/taptalk-diary/app/users/services"

	"github.com/henrioseptiano/taptalk-diary/middleware"
	"gorm.io/gorm"
)

func SwaggerRoutes(app *fiber.App) {
	route := app.Group("/swagger")
	route.Get("*", swagger.Handler)

}

func UserRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := UserRepo.New(db)
	userService := UserService.New(userRepo)

	user := &User.UserController{UserService: userService}
	r := app.Group("/api/v1")
	r.Post("/login", user.Login)
	r.Post("/register", user.Register)
	r.Get("/getcurrentdeviceid", middleware.JwtProtected(), user.GetCurrentDeviceID)
}

func DiaryRoutes(app *fiber.App, db *gorm.DB) {
	diaryRepo := DiariesRepo.New(db)
	diaryService := DiaryService.New(diaryRepo)
	diary := &Diary.DiariesController{DiariesServices: diaryService}
	r := app.Group("/api/v1/diary")
	r.Post("/create", middleware.JwtProtected(), diary.CreateDiary)
	r.Put("/update/:id", middleware.JwtProtected(), diary.UpdateDiary)
	r.Delete("/delete/:id", middleware.JwtProtected(), diary.DeleteDiary)
	r.Get("/listall", middleware.JwtProtected(), diary.ListAllDiaries)
	r.Get(":id", middleware.JwtProtected(), diary.GetDiaryById)
}
