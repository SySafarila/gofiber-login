package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var Validate *validator.Validate = validator.New()

type ValidationError struct {
	Message    string   `json:"message"`
	Errors     []string `json:"errors"`
	StatusCode int      `json:"status_code"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

func ParseErrorMessage(err error) []string {
	var errorMessages []string

	if err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			var message string
			tag := fieldError.Tag()
			fieldName := strings.ToLower(fieldError.Field())
			param := fieldError.Param()

			switch tag {
			case "required":
				message = "Field " + fieldName + " is required"
			case "email":
				message = "Field " + fieldName + " must be a valid email address"
			case "min":
				message = "Field " + fieldName + " must be at least " + param + " characters/numeric"
			case "max":
				message = "Field " + fieldName + " must be at most " + param + " characters/numeric"
			default:
				message = "Field " + fieldName + " is invalid"
			}

			errorMessages = append(errorMessages, message)
		}
	}

	return errorMessages
}
