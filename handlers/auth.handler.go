package handlers

import (
	"github.com/gofiber/fiber/v2"
	"mygo/models"
	"mygo/services"
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
	user, err := services.CheckUser(body.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Credentials not match")
	}

	// check password
	errCompare := utils.CompareHashPassword(body.Password, user.Password)
	if errCompare != nil {
		return fiber.NewError(fiber.StatusNotFound, "Credentials not match")
	}

	var userResponse models.UserResponse
	var username *string
	if user.Username.Valid {
		username = &user.Username.String
	}
	userResponse = models.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Username:  username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
		"data":    userResponse,
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
