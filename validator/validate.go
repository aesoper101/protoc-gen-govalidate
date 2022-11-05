package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func Validator() *validator.Validate {
	return validate
}

func SetDefaultValidator(v *validator.Validate) {
	validate = v
}
