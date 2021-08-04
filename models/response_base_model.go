package models

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Pagination
type Pagination struct {
	TotalRecords int         `json:"totalRecords"`
	TotalPages   int         `json:"totalPages"`
	Data         interface{} `json:"data"`
	Offset       int         `json:"offset"`
	Limit        int         `json:"Limit"`
	Page         int         `json:"page"`
	PrevPage     int         `json:"prevPage"`
	NextPage     int         `json:"nextPage"`
}

// ResponseSuccess model for swagger
type ResponseSuccess struct {
	Data struct{} `json:"data"`
}

// ResponseSuccessArray model for swagger
type ResponseSuccessArray struct {
	Data []struct{} `json:"data"`
}

// ResponseErrors model for swagger
type ResponseErrors struct {
	Error errorResponse `json:"error"`
}

// ResponseJSON ...
func ResponseJSON(c *fiber.Ctx, data interface{}) {
	c.JSON(fiber.Map{"code": http.StatusOK, "data": data})
	return
}

// ResponsePagination ...
func ResponsePagination(c *fiber.Ctx, data Pagination) {
	c.JSON(fiber.Map{"code": http.StatusOK, "data": data})
	return
}

// ResponseCreated Set response for created process
func ResponseCreated(c *fiber.Ctx, message interface{}) {
	c.JSON(fiber.Map{"code": http.StatusCreated, "message": message})
	return
}

// ResponseUpdated Set response for update process
func ResponseUpdated(c *fiber.Ctx, message interface{}) {
	c.JSON(fiber.Map{"code": http.StatusNoContent, "message": message})
	return
}

// ResponseDeleted Set response for delete process
func ResponseDeleted(c *fiber.Ctx, message interface{}) {
	if message == "" {
		message = "Resource Deleted"
	}
	c.JSON(fiber.Map{"code": http.StatusNoContent, "data": message})
	return
}

// ResponseError Set response for delete process
func ResponseError(c *fiber.Ctx, message interface{}, statusCode int) {
	c.JSON(fiber.Map{"code": statusCode, "data": message})
	return
}

// ResponseFailValidation Set response for fail validation
func ResponseFailValidation(c *fiber.Ctx, message interface{}) {
	ResponseError(c, message, 422)
	return
}

func ResponseUnauthorized(c *fiber.Ctx, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	c.JSON(fiber.Map{"code": http.StatusUnauthorized, "message": message})
	return
}

func ResponseNotFound(c *fiber.Ctx, message string) {
	if message == "" {
		message = "Resource not Found"
	}
	c.JSON(fiber.Map{"code": http.StatusUnauthorized, "message": message})
	return
}

func ResponseMethodNotAllowed(c *fiber.Ctx, message string) {
	if message == "" {
		message = "Method not Allowed"
	}
	c.JSON(fiber.Map{"code": http.StatusNotFound, "message": message})
	return
}

func ResponseRedirect(c *fiber.Ctx, url string) {
	if url == "" {
		return
	}
	c.Redirect(url, http.StatusNotFound)
	return
}
