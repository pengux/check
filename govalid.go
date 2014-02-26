package govalid

import (
	"fmt"
	"reflect"
)

// Validator is an interface for constraint types with a method of validate()
type Validator interface {
	// Validate check value against constraints
	Validate() error
}

// Validatable allows other to fetch and manipulate the data dictonary of a type
type Validatable interface {
	GetValidators() map[string][]Validator
}

// ValidationError is of type error which implement error, stringer and json.Unmarshaler interfaces
type ValidationError string

func (e ValidationError) Error() string {
	return string(e)
}
func (e ValidationError) String() string {
	return string(e)
}

// ValidationErrors is a map with string and slice of error
type ValidationErrors map[string][]ValidationError

// Validate accepts a ValidatorsGetter, fetch validators with GetValidators() and execute validation on the object
func Validate(o Validatable) (errs ValidationErrors, hasErr bool) {
	errs = make(ValidationErrors)
	for field, validators := range o.GetValidators() {
		for _, validator := range validators {
			if err := validator.Validate(); err != nil {
				if _, ok := errs[field]; !ok {
					errs[field] = make([]ValidationError, 0)
				}

				errs[field] = append(errs[field], ValidationError(err.Error()))
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

// Validate value to not be a zeroed value
func (v *NonZero) Validate() error {
	t := reflect.TypeOf(v.Value)

	switch t.Kind() {
	default:
		if reflect.DeepEqual(reflect.Zero(t).Interface(), v.Value) {
			return fmt.Errorf("require a non-zero value")
		}
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String:
		if reflect.ValueOf(v.Value).Len() == 0 {
			return fmt.Errorf("require a non-zero value")
		}
	}

	return nil
}
