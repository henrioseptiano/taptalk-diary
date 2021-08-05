package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	User "github.com/henrioseptiano/taptalk-diary/app/users/controller"
)

func SwaggerRoutes(app *fiber.App) {
	route := app.Group("/swagger")
	route.Get("*", swagger.Handler)

}

func UserRoutes(app *fiber.App) {
	user := &User.UserController{}
	r := app.Group("/api/v1")
	r.Post("/login", user.Login)
	r.Post("/register", user.Register)
	r.Get("/checkdevicelogin", user.CheckDeviceLogin)
	r.Put("/changepassword", user.ChangePassword)
}
