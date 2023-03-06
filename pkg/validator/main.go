package validator

import (
	"reflect"
	"strings"

	goValidator "github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *goValidator.Validate
}

func NewValidator() *Validator {
	validate := goValidator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	})

	return &Validator{
		validator: validate,
	}
}

func (v *Validator) Struct(s interface{}) error {
	return v.validator.Struct(s)
}
