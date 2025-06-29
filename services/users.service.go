package services

import (
	"errors"
	"mygo/database"
	"mygo/models"
)

func CheckUser(email string) (models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return user, errors.New(result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return user, errors.New(result.Error.Error())
	}

	return user, nil
}
