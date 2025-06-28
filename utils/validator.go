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
			fieldName := fieldError.Field()

			switch tag {
			case "required":
				message = "Field " + strings.ToLower(fieldName) + " is required"
			case "email":
				message = "Field " + strings.ToLower(fieldName) + " must be a valid email address"
			case "min":
				message = "Field " + strings.ToLower(fieldName) + " must be at least " + fieldError.Param() + " characters/numeric"
			case "max":
				message = "Field " + strings.ToLower(fieldName) + " must be at most " + fieldError.Param() + " characters/numeric"
			default:
				message = "Field " + strings.ToLower(fieldName) + " is invalid"
			}

			errorMessages = append(errorMessages, message)
		}
	}

	return errorMessages
}
