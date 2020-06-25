package model

import (
	"gopkg.in/go-playground/validator.v9"
)

type IceCream struct {
	ProductId             string `json:"productId" validate:"required,number"`
	Name                  string `json:"name" validate:"required"`
	ImageClosed           string `json:"imageClosed" validate:"required"`
	ImageOpen             string `json:"imageOpen" validate:"required"`
	Description           string `json:"description" validate:"required"`
	Story                 string `json:"story" validate:"required"`
	SourcingValues        []string
	Ingredients           []string
	AllergyInfo           string
	DietaryCertifications string
}

type IceCreamSearchOptions struct {
	LastId string
	Limit  int64
}

type IceCreamListResult struct {
	IceCreams []*IceCream
	LastId    string
}

func GetIceCreamValidator() *validator.Validate {
	return validator.New()
}
