// Package check allows data validation of values in different types
package check

import (
	"fmt"
	"reflect"
	"strings"
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
		"uuid":        "'%v' is an invalid uuid",
	}
)

// Validator is an interface for constraint types with a method of validate()
type Validator interface {
	// Validate check value against constraints
	Validate(v interface{}) Error
}

// Struct allows validation of structs
type Struct map[string]Validator

// Validate execute validation using the validators.
func (s Struct) Validate(v interface{}) StructError {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		panic("not a struct")
	}

	e := StructError{}
	for fieldname, validator := range s {
		field := val.FieldByName(fieldname)
		if err := validator.Validate(field.Interface()); err != nil {
			if _, ok := e[fieldname]; !ok {
				e[fieldname] = make([]Error, 0)
			}

			e[fieldname] = append(e[fieldname], err)
		}
	}

	return e
}

// Composite allows adding multiple validators to the same value
type Composite []Validator

// Validate implements Validator
func (c Composite) Validate(v interface{}) Error {
	e := ValidationError{
		errorMap: make(map[string][]interface{}),
	}
	for _, validator := range c {
		if err := validator.Validate(v); err != nil {
			errMap := err.ErrorMap()
			for msg, params := range errMap {
				e.errorMap[msg] = params
			}
		}
	}

	if len(e.errorMap) == 0 {
		return nil
	}

	return e
}

// Error is the default validation error. The Params() method returns the params
// to be used in error messages
type Error interface {
	Error() string
	ErrorMap() map[string][]interface{}
}

// ValidationError implements Error
type ValidationError struct {
	errorMap map[string][]interface{}
}

func (e ValidationError) Error() string {
	var errs []string
	for msg := range e.errorMap {
		errs = append(errs, msg)
	}

	return strings.Join(errs, "; ")
}

// ErrorMap returns the error map
func (e ValidationError) ErrorMap() map[string][]interface{} { return e.errorMap }

// NewValidationError create and return a ValidationError
func NewValidationError(key string, params ...interface{}) *ValidationError {
	return &ValidationError{
		map[string][]interface{}{
			key: params,
		},
	}
}

// StructError is a map with validation errors
type StructError map[string][]Error

// HasErrors return true if the ErrorMap has errors
func (e StructError) HasErrors() bool {
	return len(e) > 0
}

// GetErrorsByKey return all errors for a specificed key
func (e StructError) GetErrorsByKey(key string) ([]Error, bool) {
	v, ok := e[key]
	return v, ok
}

// ToMessages convert StructError to a map of field and their validation key with proper error messages
func (e StructError) ToMessages() map[string]map[string]string {
	errMessages := make(map[string]map[string]string)

	for field, validationErrors := range e {
		errMessages[field] = make(map[string]string)
		for _, err := range validationErrors {
			for key, params := range err.ErrorMap() {
				msg, ok := ErrorMessages[key]
				if !ok {
					errMessages[field][key] = "invalid data"
				} else {
					if len(params) > 0 {
						errMessages[field][key] = fmt.Sprintf(msg, params...)
					} else {
						errMessages[field][key] = msg
					}
				}
			}
		}
	}

	return errMessages
}

// NonEmpty check that the value is not a zeroed value depending on its type
type NonEmpty struct{}

// Validate value to not be a zeroed value, return error and empty slice of strings
func (validator NonEmpty) Validate(v interface{}) Error {
	t := reflect.TypeOf(v)

	switch t.Kind() {
	default:
		if reflect.DeepEqual(reflect.Zero(t).Interface(), v) {
			return NewValidationError("nonZero")
		}
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		if reflect.ValueOf(v).Len() == 0 {
			return NewValidationError("nonZero")
		}
	}

	return nil
}
