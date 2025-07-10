package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"mygo/utils"
)

type AuthUser struct {
	Id string `json:"id,omitempty"`
}

func CheckAuth(c *fiber.Ctx) error {
	token, err := utils.GetBearerToken(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	parseToken, errParse := utils.ParseToken(token)
	if errParse != nil {
		return fiber.NewError(fiber.StatusUnauthorized, errParse.Error())
	}

	user := AuthUser{Id: parseToken.UserId}

	c.Locals("user", user)
	return c.Next()
}
