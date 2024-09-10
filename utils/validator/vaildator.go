package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	alphaSpaceRegexString string = "^[a-zA-Z ]+$"
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

type CustomValidator struct {
	Validator *validator.Validate
}

func EnvValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	validValues := []string{"prod", "stage"} // Example array of valid values

	// Check if the value exists in the array
	for _, validValue := range validValues {
		if value == validValue {
			return true
		}
	}
	return false
}
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

// func New() *validator.Validate {
// 	validate := validator.New()
// 	validate.SetTagName("form")

// 	// Using the names which have been specified for JSON representations of structs, rather than normal Go field names
// 	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
// 		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
// 		if name == "-" {
// 			return ""
// 		}
// 		return name
// 	})

// 	validate.RegisterValidation("alpha_space", isAlphaSpace)

// 	return validate
// }

func ToErrResponse(err error) *ErrResponse {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ErrResponse{
			Errors: make([]string, len(fieldErrors)),
		}

		// TODO: add more errors here if needed
		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "max":
				resp.Errors[i] = fmt.Sprintf("%s must be a maximum of %s in length", err.Field(), err.Param())
			case "url":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid URL", err.Field())
			case "alpha_space":
				resp.Errors[i] = fmt.Sprintf("%s can only contain alphabetic and space characters", err.Field())
			case "datetime":
				if err.Param() == "2006-01-02" {
					resp.Errors[i] = fmt.Sprintf("%s must be a valid date", err.Field())
				} else {
					resp.Errors[i] = fmt.Sprintf("%s must follow %s format", err.Field(), err.Param())
				}
			default:
				resp.Errors[i] = fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag())
			}
		}

		return &resp
	}

	return nil
}

func isAlphaSpace(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(alphaSpaceRegexString)
	return reg.MatchString(fl.Field().String())
}
