package handlers

import (
	"github.com/gofiber/fiber/v2"
	"mygo/utils"
)

type LoginStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterStruct struct {
	FullName        string `json:"full_name" validate:"required,max=255"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8"`
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

	// check email on database
	// check password

	return c.JSON(fiber.Map{
		"message": "Login success",
		"data": fiber.Map{
			"token": "{JWT_TOKEN}",
		},
	})
}

func Register(c *fiber.Ctx) error {
	body := new(RegisterStruct)

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

	// check password and confirm password
	if body.Password != body.PasswordConfirm {
		return &utils.ValidationError{
			Message:    "Password and Confirm Passwor does not match",
			Errors:     []string{"Password and Confirm Passwor does not match"},
			StatusCode: fiber.StatusBadRequest,
		}
	}

	// check email on database
	// register credentials to database

	return c.JSON(fiber.Map{
		"message": "Register success",
	})
}

func Me(c *fiber.Ctx) error {
	_, err := utils.GetBearerToken(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Get current user",
		"data": fiber.Map{
			"id":        "{UUID}",
			"full_name": "{FULL_NAME}",
			"email":     "{EMAIL}",
		},
	})
}

func Logout(c *fiber.Ctx) error {
	_, err := utils.GetBearerToken(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Logout success",
	})
}
