package utils

import (
	"net/mail"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validatorRequest := validator.New()

	validatorRequest.RegisterValidation("password", PasswordValidator)
	validatorRequest.RegisterValidation("email", EmailValidator)

	return validatorRequest
}

func PasswordValidator(f1 validator.FieldLevel) bool {
	password := f1.Field().String()

	if len(password) < 8 {
		return false
	}

	return true
}

func EmailValidator(f1 validator.FieldLevel) bool {
	email := f1.Field().String()

	_, err := mail.ParseAddress(email)

	return err == nil
}
