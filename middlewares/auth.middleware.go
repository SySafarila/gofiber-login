package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"mygo/utils"
)

func CheckAuth(c *fiber.Ctx) error {
	token, err := utils.GetBearerToken(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	parseToken, errParse := utils.ParseToken(token)
	if errParse != nil {
		return fiber.NewError(fiber.StatusUnauthorized, errParse.Error())
	}

	userId := parseToken.UserId
	//user, errCheckUser := services.CheckUserById(userId)
	//if errCheckUser != nil {
	//	return fiber.NewError(fiber.StatusNotFound, "User not found")
	//}

	c.Locals("userId", userId)
	return c.Next()
}
