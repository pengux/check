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
	Validate() (err error, params []string)
}

// ErrorMap is a map with validation errors
type ErrorMap map[string]map[string][]string

// Add accept a key and a slice of validators which will be run and any errors
// from the validation will be saved in ErrorMap
func (e *ErrorMap) Add(key string, validators ...Validator) {
	for _, validator := range validators {
		if err, params := validator.Validate(); err != nil {
			if _, ok := (*e)[key]; !ok {
				(*e)[key] = make(map[string][]string, 1)
			}

			(*e)[key][string(err.Error())] = params
		}
	}
}

// HasErrors return true if the ErrorMap has errors
func (e *ErrorMap) HasErrors() bool {
	return len(*e) > 0
}

// GetErrorsByKey return all errors for a specificed key
func (e *ErrorMap) GetErrorsByKey(key string) (map[string][]string, bool) {
	v, ok := (*e)[key]
	return v, ok
}

// ToMessages convert ErrorMap to a map of field and their validation key with proper error messages
func (e *ErrorMap) ToMessages(messages map[string]string) map[string]map[string]string {
	errMessages := make(map[string]map[string]string)

	for field, validationErrors := range *e {
		errMessages[field] = make(map[string]string)
		for key, params := range validationErrors {
			// Convert params to []interface{}
			vals := make([]interface{}, len(params))
			for i, v := range params {
				vals[i] = v
			}

			msg, ok := ErrorMessages[key]
			if !ok {
				errMessages[field][key] = "invalid data"
			} else {
				if len(params) > 0 {
					errMessages[field][key] = fmt.Sprintf(msg, vals...)
				} else {
					errMessages[field][key] = msg
				}
			}
		}
	}

	return errMessages
}

// NonZero check that the value is not a zeroed value depending on its type
type NonZero struct {
	Value interface{}
}

// Validate value to not be a zeroed value, return error and empty slice of strings
func (v NonZero) Validate() (err error, params []string) {
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
