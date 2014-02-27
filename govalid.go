package govalid

import (
	"fmt"
	"reflect"
)

// Validator is an interface for constraint types with a method of validate()
type Validator interface {
	// Validate check value against constraints
	Validate() (err error, params []string)
}

// ValidationErrors is a map with string keys and sub maps of ValidationErrors
// as keys and slices of string with validation params
type ValidationErrors map[string]map[string][]string

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
