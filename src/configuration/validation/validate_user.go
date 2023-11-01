package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
)

var (
	Validate   = validator.New()
	translator = ut.Translator(nil)
)

func init() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLocale := en.New()
		universalTranslator := ut.New(enLocale, enLocale)
		translator, _ = universalTranslator.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(validate, translator)
	}
}

func ValidateUserError(validationError error) *resterr.RestErr {
	var jsonUnmarshallingError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonUnmarshallingError) {
		return resterr.NewBadRequestError("Invalid field type")
	} else if errors.As(validationError, &jsonValidationError) {
		errorCauses := []resterr.Causes{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: e.Translate(translator),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return resterr.NewBadRequestValidationError("There are invalid field(s)", errorCauses)
	}
	return resterr.NewBadRequestError("Error trying to convert fields")
}
