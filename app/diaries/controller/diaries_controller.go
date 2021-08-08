package controller

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/henrioseptiano/taptalk-diary/app/diaries"
	"github.com/henrioseptiano/taptalk-diary/models"
	"github.com/henrioseptiano/taptalk-diary/utils"
)

type DiariesController struct {
	DiariesServices diaries.DiariesServicesInterface
}

// Create Diary godoc
// @Summary Create Diary
// @Id CreateDiary
// @Tags Diaries
// @Security Token
// @Param create body models.ReqCreateDiary true "all fields mandatory. for datePost (ex: DD-MM-YYYY)"
// @Success 200 {object} models.ResponseSuccess "Diary Successfully Created"
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "JWT Token is not valid" "
// @Router /api/v1/diary/create [post]
func (u *DiariesController) CreateDiary(c *fiber.Ctx) error {
	modelCreateDiary := models.ReqCreateDiary{}
	if err := c.BodyParser(&modelCreateDiary); err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}
	if err := u.DiariesServices.Create(claims.ID, &modelCreateDiary); err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}
	return models.ResponseJSON(c, fiber.Map{"message": "Diary Successfully Created"})
}

// UpdateDiary godoc
// @Summary Update Diary
// @Id UpdateDiary
// @Tags Diaries
// @Security Token
// @Param id path integer true "Diary ID"
// @Param update body models.ReqUpdateDiary true "all fields mandatory. for datePost (ex: DD-MM-YYYY)"
// @Success 200 {object} models.ResponseSuccess "Diary Successfully Updated"
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Invalid ID" "
// @Router /api/v1/diary/update/{id} [put]
func (u *DiariesController) UpdateDiary(c *fiber.Ctx) error {
	diaryID, err := c.ParamsInt("id")
	if err != nil {
		return models.ResponseError(c, "Invalid ID", 401)
	}
	modelUpdateDiary := models.ReqUpdateDiary{}
	if err := c.BodyParser(&modelUpdateDiary); err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}
	if err := u.DiariesServices.Update(claims.ID, int64(diaryID), &modelUpdateDiary); err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}
	return models.ResponseJSON(c, fiber.Map{"message": "Diary Successfully Updated"})
}

// DeleteDiary godoc
// @Summary Delete Diary
// @Id DeleteDiary
// @Tags Diaries
// @Security Token
// @Param id path integer true "Diary ID"
// @Success 200 {object} models.ResponseSuccess "Diary Successfully Deleted" "
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/diary/delete/{id} [delete]
func (u *DiariesController) DeleteDiary(c *fiber.Ctx) error {
	diaryID, err := c.ParamsInt("id")
	if err != nil {
		return models.ResponseError(c, "Invalid ID", 401)
	}
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}
	u.DiariesServices.Delete(claims.ID, int64(diaryID))
	return models.ResponseJSON(c, fiber.Map{"token": "Diary Successfully Deleted"})
}

// ListAllDiaries godoc
// @Summary List All Diaries
// @Id ListAllDiaries
// @Tags Diaries
// @Security Token
// @Param page query string false " "
// @Param limit query string false " "
// @Param year query string true "year of posted diaries"
// @param quarter query string true "Quarter of diaries: ex (1 : January - March, 2: April - June, 3: July - September, 4: October - December)"
// @Success 200 {object} models.ResponseSuccess "models.Pagination"
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/diary/listall [get]
func (u *DiariesController) ListAllDiaries(c *fiber.Ctx) error {
	/*This is for default value*/
	yearNow, monthNow, _ := time.Now().Date()
	intMonthNow := int(monthNow)
	defaultQuarter := utils.GetQuarter(int64(intMonthNow))
	/*Get the input Query*/
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return models.ResponseError(c, "Page must be numeric", 500)
	}
	limit, err := strconv.Atoi(c.Query("limit", "1"))
	if err != nil {
		return models.ResponseError(c, "limit must be numeric", 500)
	}
	year := c.Query("year", strconv.Itoa(yearNow))
	if checkIsNumeric := utils.IsNumeric(year); checkIsNumeric == false {
		if err != nil {
			return models.ResponseError(c, "year must be numeric", 500)
		}
	}
	quarter := c.Query("quarter", defaultQuarter)
	if checkIsNumeric := utils.IsNumeric(quarter); checkIsNumeric == false {
		if err != nil {
			return models.ResponseError(c, "quarter must be numeric", 500)
		}
	}
	monthRange := utils.GetMonthRageFromQuarter(quarter)

	return models.ResponsePagination(c, *u.DiariesServices.GetAllDiariesPagination(year, monthRange, page, limit))
}

// GetDiaryById godoc
// @Summary Get Diary By Id
// @Id GetDiaryById
// @Tags Diaries
// @Security Token
// @Success 200 {object} entity.UserDiary "entity.UserDiary"
// @Failure 422 {object} models.ResponseErrors "code: 422, message: "Invalid request" "
// @Failure 401 {object} models.ResponseErrors "code: 401, message: "Username or password not valid, please try again" "
// @Router /api/v1/diary/{id} [get]
func (u *DiariesController) GetDiaryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return models.ResponseError(c, "ID must be numeric", 500)
	}
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return models.ResponseError(c, err.Error(), 500)
	}
	return models.ResponseJSON(c, fiber.Map{"data": u.DiariesServices.GetDiariesByID(int64(id), claims.ID)})
}
