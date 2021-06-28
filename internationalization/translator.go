package internationalization

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/wevr-in/common/internationalization/locale"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Translator(l locale.Locale) (ut.Translator, *validator.Validate) {
	var loc locales.Translator
	switch l {
	case locale.EN:
		loc = en.New()
		break
	default:
		loc = en.New()
		break
	}
	uni = ut.New(loc, loc)

	trans, _ := uni.GetTranslator(string(l))

	validate = validator.New()
	err := enTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, nil
	}
	translateOverride(trans)
	return trans, validate
}

func translateOverride(trans ut.Translator) {
	re := validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field.", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
	if re != nil {
		return
	}

	ve := validate.RegisterTranslation("alpha", trans, func(ut ut.Translator) error {
		return ut.Add("alpha", "{0} of type string is required.", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("alpha", fe.Field())

		return t
	})
	if ve != nil {
		return
	}
}
