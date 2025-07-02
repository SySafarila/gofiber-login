package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserTokenStruct struct {
	UserId string `json:"user_id"`
	Iat    int64  `json:"iat"`
	Exp    int64  `json:"exp"`
}

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

func CreateUserToken(data UserTokenStruct) (string, error) {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte("MYSECRETKEY")
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": data.UserId,
		"iat":     data.Iat,
		"exp":     data.Exp,
	})
	s, _ = t.SignedString(key)

	return s, nil
}
