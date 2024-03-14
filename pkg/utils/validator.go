package utils

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func ValidateStruct(req interface{}) error {
	validate = validator.New()
	return validate.Struct(req)
}

func translateOverride() ut.Translator {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "field {0} cannot be empty!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
	return trans
}

func GetValidatorFieldError(err error) map[string]string {
	translator := translateOverride()
	result := map[string]string{}
	if _, ok := err.(*validator.InvalidValidationError); ok {
		result["interface"] = err.Error()
	}
	if val, ok := err.(validator.ValidationErrors); ok {
		for _, err := range val {
			result[err.Field()] = err.Translate(translator)
		}
		return result
	}
	return nil
}
