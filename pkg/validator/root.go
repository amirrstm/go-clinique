package validator

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		errMsg := fmt.Sprintf("validation failed on '%s'", err.Tag())
		param := err.Param()
		if param != "" {
			errMsg = fmt.Sprintf("%s. allow: %s", errMsg, param)
		}
		fields[err.Field()] = errMsg
	}

	return fields
}
