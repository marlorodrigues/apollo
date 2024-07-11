package authentication

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type sLogin struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:required" json:"password"`
}

func validLogin(data *sLogin) string {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		return _formatValidationError(validationErrors)
	}

	return ""
}

func _formatValidationError(err validator.ValidationErrors) string {
	var errorMsg string

	for _, e := range err {
		errorMsg += fmt.Sprintf("Campo '%s' falhou na validação com tag '%s'\n", e.Field(), e.Tag())
	}

	return errorMsg
}
