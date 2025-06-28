package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"mygo/utils"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// default status code
	code := fiber.StatusInternalServerError
	message := "Internal server error"

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
		message = fiberErr.Message
	}

	var validationErr *utils.ValidationError
	if errors.As(err, &validationErr) {
		ctx.Status(validationErr.StatusCode)
		return ctx.JSON(validationErr)
	}

	if err != nil {
		if code == fiber.StatusInternalServerError {
			log.Error(err, err.Error())
		}

		ctx.Status(code)
		return ctx.JSON(fiber.Map{
			"message":     message,
			"status_code": code,
		})
	}

	return nil
}
