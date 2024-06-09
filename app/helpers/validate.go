package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}
