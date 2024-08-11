package val_impls

import (
	"errors"
	"github.com/lysand-org/versia-go/internal/validators"
	"reflect"
	"strings"

	"github.com/go-logr/logr"
	en_locale "github.com/go-playground/locales/en"
	universal_translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/lysand-org/versia-go/ent/schema"
	"github.com/lysand-org/versia-go/internal/api_schema"
)

var _ validators.BodyValidator = (*BodyValidatorImpl)(nil)

type BodyValidatorImpl struct {
	validator    *validator.Validate
	translator   *universal_translator.UniversalTranslator
	enTranslator universal_translator.Translator

	log logr.Logger
}

func NewBodyValidator(log logr.Logger) *BodyValidatorImpl {
	en := en_locale.New()
	translator := universal_translator.New(en, en)
	trans, ok := translator.GetTranslator("en")
	if !ok {
		panic("failed to get \"en\" translator")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		panic("failed to register default translations")
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.RegisterValidation("username_regex", func(fl validator.FieldLevel) bool {
		return schema.ValidateUsername(fl.Field().String()) == nil
	}); err != nil {
		panic("failed to register username_regex validator")
	}

	if err := validate.RegisterTranslation("username_regex", trans, func(ut universal_translator.Translator) error {
		return trans.Add("user_regex", "{0} must match '^[a-z0-9_-]+$'!", true)
	}, func(ut universal_translator.Translator, fe validator.FieldError) string {
		t, _ := ut.T("user_regex", fe.Field())
		return t
	}); err != nil {
		panic("failed to register user_regex translation")
	}

	return &BodyValidatorImpl{
		validator:    validate,
		translator:   translator,
		enTranslator: trans,
	}
}

func (i BodyValidatorImpl) Validate(v any) error {
	err := i.validator.Struct(v)
	if err == nil {
		return nil
	}

	var invalidValidationError *validator.InvalidValidationError
	if errors.As(err, &invalidValidationError) {
		panic(invalidValidationError)
	}

	i.log.Error(err, "Failed to validate object")

	errs := make([]string, 0)
	for _, err := range err.(validator.ValidationErrors) {
		errs = append(errs, err.Translate(i.enTranslator))
	}

	return api_schema.ErrBadRequest(map[string]any{"validation": errs})
}
