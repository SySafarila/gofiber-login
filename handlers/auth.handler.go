package handlers

import (
	"github.com/gofiber/fiber/v2"
	"mygo/utils"
)

type LoginStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

var validate = utils.Validate

func Login(c *fiber.Ctx) error {
	body := new(LoginStruct)

	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	errValidate := validate.Struct(body)
	if errValidate != nil {
		errMessages := utils.ParseErrorMessage(errValidate)

		return &utils.ValidationError{
			Message:    errMessages[0],
			Errors:     errMessages,
			StatusCode: fiber.StatusBadRequest,
		}
	}

	return c.JSON(fiber.Map{
		"message": "Login",
		"data": fiber.Map{
			"email":    body.Email,
			"password": body.Password,
		},
	})
}
