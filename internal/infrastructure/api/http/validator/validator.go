package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// wrapper implementation.
type wrapper struct {
	validator *validator.Validate
}

func New() echo.Validator {
	return &wrapper{
		validator: validator.New(),
	}
}

// Validate data
func (v *wrapper) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
