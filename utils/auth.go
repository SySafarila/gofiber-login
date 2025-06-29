package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetBearerToken(c *fiber.Ctx) (string, error) {
	headers := c.GetReqHeaders()
	bearerToken := headers["Authorization"]

	if len(bearerToken) < 1 || bearerToken[0] == "" {
		return "", errors.New("bearer token required")
	}

	return bearerToken[0], nil
}

func CompareHashPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
