package validation

import (
	"github.com/deall-users/pkg/utils"
	"github.com/go-playground/validator/v10"
)

func ValidateDecryptText(fl validator.FieldLevel) bool {
	if words, ok := fl.Field().Interface().(string); ok {
		_, err := utils.Decrypt(words)
		return err == nil
	}

	return false
}
