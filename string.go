package check

import (
	"regexp"
	"strconv"
	"strings"
)

// MinChar validates that a string must have a length minimum of its constraint
type MinChar struct {
	Constraint int
	Value      string
}

// Validate check value against constraint
func (v MinChar) Validate() Error {
	if len(v.Value) < v.Constraint {
		return &ValidationError{"minChar", []interface{}{strconv.Itoa(v.Constraint)}}
	}

	return nil
}

// MaxChar validates that a string must have a length maximum of its constraint
type MaxChar struct {
	Constraint int
	Value      string
}

// Validate check value against constraint
func (v MaxChar) Validate() Error {
	if len(v.Value) > v.Constraint {
		return &ValidationError{"maxChar", []interface{}{strconv.Itoa(v.Constraint)}}
	}

	return nil
}

// Email is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
type Email struct {
	Value string
}

// Validate email addresses
func (v Email) Validate() Error {
	if !strings.Contains(v.Value, "@") || string(v.Value[0]) == "@" || string(v.Value[len(v.Value)-1]) == "@" {
		return &ValidationError{"email", []interface{}{v.Value}}
	}

	return nil
}

// Regex allow validation usig regular expressions
type Regex struct {
	Constraint string
	Value      string
}

// Validate using regex
func (v Regex) Validate() Error {
	regex, err := regexp.Compile(v.Constraint)
	if err != nil {
		return &ValidationError{err.Error(), nil}
	}

	if !regex.MatchString(v.Value) {
		return &ValidationError{"regex", []interface{}{v.Value, v.Constraint}}
	}

	return nil
}
