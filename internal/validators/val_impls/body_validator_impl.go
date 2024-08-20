package val_impls

import (
	"errors"
	"github.com/lysand-org/versia-go/ent/schema"
	"github.com/lysand-org/versia-go/internal/validators"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-logr/logr"
	en_locale "github.com/go-playground/locales/en"
	universal_translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/lysand-org/versia-go/internal/api_schema"
)

type bodyValidator struct {
	Translated string
	Validate   func(fl validator.FieldLevel) bool
}

var (
	_ validators.BodyValidator = (*BodyValidatorImpl)(nil)

	fullUserRegex = regexp.MustCompile("^@([a-z0-9_-]+)(?:@([a-zA-Z0-1-_.]+\\.[a-zA-Z0-9-z]+(?::[0-9]+)?))?$")
	domainRegex   = regexp.MustCompile("^[a-zA-Z0-9-_.]+.[a-zA-Z0-9-z]+(?::[0-9]+)?$")
)

type BodyValidatorImpl struct {
	validator    *validator.Validate
	translator   *universal_translator.UniversalTranslator
	enTranslator universal_translator.Translator

	log logr.Logger
}

func NewBodyValidator(log logr.Logger) *BodyValidatorImpl {
	en := en_locale.New()
	translator := universal_translator.New(en, en)
	enTranslator, ok := translator.GetTranslator("en")
	if !ok {
		panic("failed to get \"en\" translator")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := en_translations.RegisterDefaultTranslations(validate, enTranslator); err != nil {
		panic("failed to register default translations")
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	bodyValidators := map[string]bodyValidator{
		"username_regex": {
			Translated: "{0} must match '^[a-z0-9_-]+$'!",
			Validate: func(fl validator.FieldLevel) bool {
				return schema.ValidateUsername(fl.Field().String()) == nil
			},
		},
		"full_user_regex": {
			Translated: "{0} must match '^@[a-z0-9_-]+$' or '^@[a-z0-9_-]+@[a-zA-Z0-1-_.]+\\.[a-zA-Z0-9-z]+(?::[0-9]+)?))?$'",
			Validate: func(fl validator.FieldLevel) bool {
				f := fl.Field()
				if f.Type().String() == "string" {
					return fullUserRegex.Match([]byte(f.String()))
				}

				return false
			},
		},
		"domain_regex": {
			Translated: "{0} must match '^[a-zA-Z0-9-_.]+.[a-zA-Z0-9-z]+(?::[0-9]+)?$'",
			Validate: func(fl validator.FieldLevel) bool {
				f := fl.Field()
				t := f.Type().String()
				if t == "string" {
					return domainRegex.Match([]byte(f.String()))
				}

				log.V(-1).Info("got wrong type: %s\n", t)

				return false
			},
		},
	}

	for identifier, v := range bodyValidators {
		if err := validate.RegisterValidation(identifier, v.Validate); err != nil {
			log.Error(err, "failed to register validator", "identifier", identifier)
		}

		register := func(ut universal_translator.Translator) error {
			return enTranslator.Add(identifier, v.Translated, true)
		}

		translate := func(ut universal_translator.Translator, fe validator.FieldError) string {
			t, _ := ut.T(identifier, fe.Field())
			return t
		}

		if err := validate.RegisterTranslation(identifier, enTranslator, register, translate); err != nil {
			log.Error(err, "failed to register validator translator", "identifier", identifier)
		}
	}

	return &BodyValidatorImpl{
		validator:    validate,
		translator:   translator,
		enTranslator: enTranslator,

		log: log,
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
