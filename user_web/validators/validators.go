package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok , _ := regexp.MatchString(`^\d{10}$`, mobile)
	if !ok{
		return false
	}
	return true
}

func ValidateEmail(fl validator.FieldLevel) bool{
	email := fl.Field().String()
	ok , _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email)
	if !ok{
		return false
	}
	return true
}
