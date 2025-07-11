package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"mygo/database"
	"mygo/models"
	"strings"
	"time"
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

	// split bearer token
	results := strings.Split(bearerToken[0], "Bearer ")

	return results[1], nil
}

func CompareHashPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func ParseToken(token string) (*models.TokenClaims, error) {
	claims := &models.TokenClaims{}
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("MYSECRETKEY"), nil
	})

	if err != nil {
		fmt.Printf("Token parse error: %v\n", err)
		return nil, err
	}

	if !tokenClaims.Valid {
		fmt.Printf("Token invalid")
		return nil, err
	}

	return claims, nil
}

func CacheToken(token string) error {
	if !database.IsRedisConnected {
		return nil
	}

	err := database.Redis.Set(database.Ctx, token, token, time.Hour*24).Err()
	return err
}
