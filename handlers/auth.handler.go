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

func Logout(c *fiber.Ctx) error {
	// get token from authorization
	headers := c.GetReqHeaders()
	bearerToken := headers["Authorization"]

	if len(bearerToken) < 1 || bearerToken[0] == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Bearer token required")
	}

	return c.JSON(fiber.Map{
		"message": "Logout success",
	})
}
