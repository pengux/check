// Package check allows data validation of values in different types
package check

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
		"email":       "'%v' is an invalid email address",
		"regex":       "'%v' does not match '%v'",
	}
)

// Validator is an interface for constraint types with a method of validate()
type Validator interface {
	// Validate check value against constraints
	Validate() Error
}

// Error is the default validation error. The Params() method returns the params
// to be used in error messages
type Error interface {
	Error() string
	Params() []interface{}
}

// ValidationError is an error implementation that includes error params
// for usage in error messages
type ValidationError struct {
	message string
	params  []interface{}
}

func (e *ValidationError) Error() string { return e.message }

// Params return the error parameters to be used in error messages
func (e *ValidationError) Params() []interface{} { return e.params }

// ErrorMap is a map with validation errors
type ErrorMap map[string][]Error

// Add accept a key and a slice of validators which will be run and any errors
// from the validation will be saved in ErrorMap
func (e *ErrorMap) Add(key string, validators ...Validator) {
	for _, validator := range validators {
		if err := validator.Validate(); err != nil {
			if _, ok := (*e)[key]; !ok {
				(*e)[key] = make([]Error, 0)
			}

			(*e)[key] = append((*e)[key], err)
		}
	}
}

// HasErrors return true if the ErrorMap has errors
func (e *ErrorMap) HasErrors() bool {
	return len(*e) > 0
}

// GetErrorsByKey return all errors for a specificed key
func (e *ErrorMap) GetErrorsByKey(key string) ([]Error, bool) {
	v, ok := (*e)[key]
	return v, ok
}

// ToMessages convert ErrorMap to a map of field and their validation key with proper error messages
func (e *ErrorMap) ToMessages(messages map[string]string) map[string]map[string]string {
	errMessages := make(map[string]map[string]string)

	for field, validationErrors := range *e {
		errMessages[field] = make(map[string]string)
		for _, err := range validationErrors {
			key := err.Error()
			msg, ok := ErrorMessages[err.Error()]
			if !ok {
				errMessages[field][key] = "invalid data"
			} else {
				if len(err.Params()) > 0 {
					errMessages[field][key] = fmt.Sprintf(msg, err.Params()...)
				} else {
					errMessages[field][key] = msg
				}
			}
		}
	}

	return errMessages
}

// NonEmpty check that the value is not a zeroed value depending on its type
type NonEmpty struct {
	Value interface{}
}

// Validate value to not be a zeroed value, return error and empty slice of strings
func (v NonEmpty) Validate() Error {
	t := reflect.TypeOf(v.Value)

	switch t.Kind() {
	default:
		if reflect.DeepEqual(reflect.Zero(t).Interface(), v.Value) {
			return &ValidationError{"nonZero", nil}
		}
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		if reflect.ValueOf(v.Value).Len() == 0 {
			return &ValidationError{"nonZero", nil}
		}
	}

	return nil
}
