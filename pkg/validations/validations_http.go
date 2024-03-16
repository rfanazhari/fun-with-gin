package validations

import (
	"github.com/gin-gonic/gin/binding"
	validator2 "github.com/go-playground/validator/v10"
)

func MaxLengthWhenNotEmpty(fl validator2.FieldLevel) bool {
	title := fl.Field().String()
	if title != "" && len(title) > 255 {
		return false
	}

	status := fl.Field().String()
	if status != "" && len(status) > 50 {
		return false
	}
	return true
}

func RegisterBindingValidation() {
	if v, ok := binding.Validator.Engine().(*validator2.Validate); ok {
		_ = v.RegisterValidation("maxLengthWhenNotEmpty", MaxLengthWhenNotEmpty)
	}
}
