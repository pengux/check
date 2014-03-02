package check

import (
	"regexp"
	"strings"
)

// MinChar validates that a string must have a length minimum of its constraint
type MinChar struct {
	Constraint int
}

// Validate check value against constraint
func (validator MinChar) Validate(v interface{}) Error {
	if len(v.(string)) < validator.Constraint {
		return NewValidationError("minChar", validator.Constraint)
	}

	return nil
}

// MaxChar validates that a string must have a length maximum of its constraint
type MaxChar struct {
	Constraint int
}

// Validate check value against constraint
func (validator MaxChar) Validate(v interface{}) Error {
	if len(v.(string)) > validator.Constraint {
		return NewValidationError("maxChar", validator.Constraint)
	}

	return nil
}

// Email is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
type Email struct {
}

// Validate email addresses
func (validator Email) Validate(v interface{}) Error {
	if !strings.Contains(v.(string), "@") || string(v.(string)[0]) == "@" || string(v.(string)[len(v.(string))-1]) == "@" {
		return NewValidationError("email", v)
	}

	return nil
}

// Regex allow validation usig regular expressions
type Regex struct {
	Constraint string
}

// Validate using regex
func (validator Regex) Validate(v interface{}) Error {
	regex, err := regexp.Compile(validator.Constraint)
	if err != nil {
		panic(err)
	}

	if !regex.MatchString(v.(string)) {
		return NewValidationError("regex", v, validator.Constraint)
	}

	return nil
}
