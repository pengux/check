package govalid

import (
	"fmt"
	"reflect"
)

var (
	// ErrorMessages contains default error messages
	ErrorMessages = map[string]string{
		"nonZero":     "value cannot be empty",
		"before":      "%v is not before %v",
		"after":       "%v is not after %v",
		"lowerThan":   "%v is not lower than %v",
		"greaterThan": "%v is not greater than %v",
		"minChar":     "too short, minimum %v characters",
		"maxChar":     "too long, minimum %v characters",
		"email":       "%v is an invalid email address",
		"regex":       "%v does not match %v",
	}
)

// Validator is an interface for constraint types with a method of validate()
type Validator interface {
	// Validate check value against constraints
	Validate() (err error, params []string)
}

// ValidationErrors is a map with string keys and sub maps of ValidationErrors
// as keys and slices of string with validation params
type ValidationErrors map[string]map[string][]string

// ToMessages convert ValidationErrors to a map of field and their validation key with proper error messages
func (v ValidationErrors) ToMessages(messages map[string]string) map[string]map[string]string {
	errMessages := make(map[string]map[string]string)

	for field, validationErrors := range v {
		errMessages[field] = make(map[string]string)
		for key, params := range validationErrors {
			msg, ok := ErrorMessages[key]
			if !ok {
				errMessages[field][key] = "invalid data"
			} else {
				if len(params) > 0 {
					errMessages[field][key] = fmt.Sprintf(msg, params)
				} else {
					errMessages[field][key] = msg
				}
			}
		}
	}

	return errMessages
}

// Validate accepts a map of string as keys and slice of Validator structs,
// execute validation with the Validator and return any errors
func Validate(m map[string][]Validator) (errs ValidationErrors, hasErr bool) {
	errs = make(ValidationErrors)
	for field, validators := range m {
		for _, validator := range validators {
			if err, params := validator.Validate(); err != nil {
				if _, ok := errs[field]; !ok {
					errs[field] = make(map[string][]string, 1)
				}

				errs[field][string(err.Error())] = params
				hasErr = true
			}
		}
	}

	return errs, hasErr
}

// NonZero check that the value is not a zeroed value depending on its type
type NonZero struct {
	Value interface{}
}

// Validate value to not be a zeroed value, return error and empty slice of strings
func (v *NonZero) Validate() (err error, params []string) {
	t := reflect.TypeOf(v.Value)

	switch t.Kind() {
	default:
		if reflect.DeepEqual(reflect.Zero(t).Interface(), v.Value) {
			return fmt.Errorf("nonZero"), params
		}
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		if reflect.ValueOf(v.Value).Len() == 0 {
			return fmt.Errorf("nonZero"), params
		}
	}

	return nil, params
}
