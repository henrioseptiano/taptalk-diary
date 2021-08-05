package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henrioseptiano/taptalk-diary/app/diaries"
	"github.com/henrioseptiano/taptalk-diary/models"
)

type DiariesController struct {
	DiariesServices diaries.DiariesServicesInterface
}

// Create Diary godoc
// @Summary Create Diary
// @Id CreateDiary
// @Tags Diaries
// @Security token
// @Success 200 {object} models.ResponseSuccess "token: "exampletokenresponse" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/diary/create [post]
func (u *DiariesController) CreateDiary(c *fiber.Ctx) error {
	return models.ResponseJSON(c, fiber.Map{"token": "createDiary"})
}

// UpdateDiary godoc
// @Summary Update Diary
// @Id UpdateDiary
// @Tags Diaries
// @Security token
// @Success 200 {object} models.ResponseSuccess "token: "exampletokenresponse" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/diary/update/{id} [put]
func (u *DiariesController) UpdateDiary(c *fiber.Ctx) error {
	return models.ResponseJSON(c, fiber.Map{"token": "updateDiary"})
}

// DeleteDiary godoc
// @Summary Delete Diary
// @Id DeleteDiary
// @Tags Diaries
// @Security token
// @Success 200 {object} models.ResponseSuccess "token: "exampletokenresponse" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/diary/delete/{id} [delete]
func (u *DiariesController) DeleteDiary(c *fiber.Ctx) error {
	return models.ResponseJSON(c, fiber.Map{"token": "delete Diary"})
}

// ListAllDiaries godoc
// @Summary List All Diaries
// @Id ListAllDiaries
// @Tags Diaries
// @Security token
// @Success 200 {object} models.ResponseSuccess "token: "exampletokenresponse" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/diary/listall [get]
func (u *DiariesController) ListAllDiaries(c *fiber.Ctx) error {
	return models.ResponseJSON(c, fiber.Map{"token": "List All Diaries"})
}

// GetDiaryById godoc
// @Summary Get Diary By Id
// @Id GetDiaryById
// @Tags Diaries
// @Security token
// @Success 200 {object} models.ResponseSuccess "token: "exampletokenresponse" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/diary/{id} [get]
func (u *DiariesController) GetDiaryById(c *fiber.Ctx) error {
	return models.ResponseJSON(c, fiber.Map{"token": "Get Diary By Id"})
}
