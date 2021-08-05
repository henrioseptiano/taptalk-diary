package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henrioseptiano/taptalk-diary/app/users"
	"github.com/henrioseptiano/taptalk-diary/models"
)

type UserController struct {
	UserService users.UserServicesInterfaces
}

// Login godoc
// @Summary Login User
// @Id LoginUser
// @Tags User
// @Param login body models.ReqUserLogin true "all fields mandatory"
// @Success 200 {object} models.ResponseSuccess "token: "exampletokenresponse" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/login [post]
func (u *UserController) Login(c *fiber.Ctx) error {
	req := models.ReqUserLogin{}
	if err := c.BodyParser(&req); err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}
	tokenString, err := u.UserService.LoginUser(&req)
	if err != nil {
		return models.ResponseError(c, "Cannot Create Token", 401)
	}
	return models.ResponseJSON(c, fiber.Map{"token": tokenString})
}

// Register godoc
// @Summary Register User
// @Id RegisterUser
// @Tags User
// @Param register body models.ReqUserRegister true "all fields mandatory"
// @Success 200 {object} models.ResponseSuccess "token: "exampletokenresponse" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/register [post]
func (u *UserController) Register(c *fiber.Ctx) error {
	return models.ResponseJSON(c, fiber.Map{"token": "register"})
}

// CheckDeviceLogin godoc
// @Summary Check Device Login
// @Id CheckDeviceLogin
// @Tags User
// @Security token
// @Success 200 {object} models.ResponseSuccess "token: "exampletokenresponse" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/checkdevicelogin [get]
func (u *UserController) CheckDeviceLogin(c *fiber.Ctx) error {
	return models.ResponseJSON(c, fiber.Map{"token": "checkDeviceLogin"})
}

// ChangePassword godoc
// @Summary Change Password
// @Id ChangePassword
// @Tags User
// @Security token
// @Success 200 {object} models.ResponseSuccess "token: "exampletokenresponse" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/changepassword [put]
func (u *UserController) ChangePassword(c *fiber.Ctx) error {
	return models.ResponseJSON(c, fiber.Map{"token": "changePassword"})
}
