package services

import (
	"errors"
	"fmt"
	"mygo/database"
	"mygo/models"
)

func CheckUserByEmail(email string) (models.User, error) {
	var user models.User
	result := database.DB.Table("users").Where("email = ?", email).First(&user)

	if result.Error != nil {
		fmt.Println("error:", result.Error)
		//return user, errors.New(result.Error.Error())
	}

	if result.RowsAffected == 0 {
		fmt.Println("rows affected 0")
		return user, errors.New(result.Error.Error())
	}

	return user, nil
}

func CheckUserById(id string) (models.User, error) {
	var user models.User
	result := database.DB.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		fmt.Println("error:", result.Error)
		//return user, errors.New(result.Error.Error())
	}

	if result.RowsAffected == 0 {
		fmt.Println("rows affected 0")
		return user, errors.New(result.Error.Error())
	}

	return user, nil
}

func RegisterUser(user models.User) {
	process := database.DB.Table("users").Create(&user)

	fmt.Println(process.Error)
	fmt.Println(process.RowsAffected)

	if process.Error != nil {
		fmt.Println("error:", process.Error)
	}

	if process.RowsAffected == 0 {
		fmt.Println("rows affected 0")
	}
}
