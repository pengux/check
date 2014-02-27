package govalid

import (
	"fmt"
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
func (v MinChar) Validate() (err error, params []string) {
	if len(v.Value) < v.Constraint {
		return fmt.Errorf("minChar"), []string{strconv.Itoa(v.Constraint)}
	}

	return nil, params
}

// MaxChar validates that a string must have a length maximum of its constraint
type MaxChar struct {
	Constraint int
	Value      string
}

// Validate check value against constraint
func (v MaxChar) Validate() (err error, params []string) {
	if len(v.Value) > v.Constraint {
		return fmt.Errorf("maxChar"), []string{strconv.Itoa(v.Constraint)}
	}

	return nil, params
}

// Email is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
type Email struct {
	Value string
}

// Validate email addresses
func (v Email) Validate() (err error, params []string) {
	if !strings.Contains(v.Value, "@") || string(v.Value[0]) == "@" || string(v.Value[len(v.Value)-1]) == "@" {
		return fmt.Errorf("email"), []string{v.Value}
	}

	return nil, params
}

// Regex allow validation usig regular expressions
type Regex struct {
	Constraint string
	Value      string
}

// Validate using regex
func (v Regex) Validate() (err error, params []string) {
	regex, err := regexp.Compile(v.Constraint)
	if err != nil {
		return err, params
	}

	if !regex.MatchString(v.Value) {
		return fmt.Errorf("regex"), []string{v.Value, v.Constraint}
	}

	return nil, params
}
