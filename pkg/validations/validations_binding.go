package validations

import (
	"github.com/go-playground/validator/v10"
	"strconv"
)

type CustomValidation struct {
	DescriptionMaxLen int
}

func (cv *CustomValidation) ValidateStatus(fl validator.FieldLevel) bool {
	// Retrieve TitleMaxLen from the field tag
	field, _ := fl.Parent().Type().FieldByName(fl.FieldName())
	maxLenTag := field.Tag.Get("statusMaxLen")
	if maxLenTag == "" {
		return false // Return false if titleMaxLen tag is not provided
	}
	statusMaxVal, err := strconv.Atoi(maxLenTag)
	if err != nil {
		return false // Return false if titleMaxLen cannot be parsed
	}

	fieldValue := fl.Field().String()
	if fieldValue != "" && len(fieldValue) <= statusMaxVal {
		return false
	}
	return true
}

func RegisterCustomValidation(v *validator.Validate, cv *CustomValidation) error {
	v.RegisterValidation("statusMaxLen", cv.ValidateStatus)
	return nil
}

// CustomBinding performs validation on the provided data using the registered validator instance
func CustomBinding(v *validator.Validate, data interface{}, customValidation *CustomValidation) error {
	if err := RegisterCustomValidation(v, customValidation); err != nil {
		return err
	}

	if err := v.Struct(data); err != nil {
		return err
	}

	return nil
}
