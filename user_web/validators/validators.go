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
