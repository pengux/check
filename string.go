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
// and that it is not in the first or last character of the string. Also if there is only one "@" and if there is a "." in the string after the "@".
type Email struct{}

// Validate email addresses
func (validator Email) Validate(v interface{}) Error {
	err := NewValidationError("email", v)

	s, ok := v.(string)
	if !ok {
		return err
	}

	email := []rune(s)
	at := '@'
	dot := '.'
	firstChar := email[0]
	lastChar := email[len(email)-1]
	emailParts := strings.Split(string(s), "@")

	if len(emailParts) != 2 || // not containing "@"
		firstChar == at || // "@bar"
		lastChar == at || // "foo@"
		lastChar == dot || // "foo@bar."
		!strings.ContainsRune(emailParts[1], dot) || // "foo@bar"
		[]rune(emailParts[1])[0] == dot { // "foo@.bar"
		return err
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

// UUID verify a string in the UUID format xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx
type UUID struct{}

// Validate checks a string as correct UUID format
func (validator UUID) Validate(v interface{}) Error {
	regex := regexp.MustCompile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")

	if !regex.MatchString(v.(string)) {
		return NewValidationError("uuid", v)
	}

	return nil
}
