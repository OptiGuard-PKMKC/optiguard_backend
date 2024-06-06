package helpers

import "github.com/go-playground/validator/v10"

func GetValidationErrors(err error) map[string]string {
	validationErrors := err.(validator.ValidationErrors)
	errorsMap := make(map[string]string)
	for _, err := range validationErrors {
		errorsMap[err.Field()] = err.Tag()
	}

	return errorsMap
}
