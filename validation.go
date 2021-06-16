package common

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
)

type ErrorTranslator struct {
	ENTranslator ut.Translator
	IDTranslator ut.Translator
}

func NewDefaultTranslator(v *validator.Validate) *ErrorTranslator {
	english := en.New()
	uni := ut.New(english, english)
	enTrans, _ := uni.GetTranslator("en")
	enTranslations.RegisterDefaultTranslations(v, enTrans)

	id := id.New()
	uni = ut.New(id, id)
	idTrans, _ := uni.GetTranslator("id")
	idTranslations.RegisterDefaultTranslations(v, idTrans)

	return &ErrorTranslator{
		ENTranslator: enTrans,
		IDTranslator: idTrans,
	}
}

func translateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Sprint(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

type Validator struct {
	Validator  *validator.Validate
	Translator *ErrorTranslator
}

func NewValidator(validator *validator.Validate, trans *ErrorTranslator) *Validator {
	return &Validator{
		Validator:  validator,
		Translator: trans,
	}
}

func (v *Validator) Validate(op string, i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {

		return NewValidationException(
			op,
			translateError(err, v.Translator.ENTranslator),
		)
	}
	return nil
}
