package model

import "gopkg.in/go-playground/validator.v9"

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"email" validate:"required,min=6"`
}

func GetUserValidator() *validator.Validate {
	return validator.New()
}
