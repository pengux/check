package govalid

import (
	"fmt"
	"regexp"
	"strings"
)

// MinChar validates that a string must have a length minimum of its constraint
type MinChar struct {
	Constraint int
	Value      string
}

// Validate check value against constraint
func (v *MinChar) Validate() error {
	if len(v.Value) < v.Constraint {
		return fmt.Errorf("too short, must be at least %v characters", v.Value)
	}

	return nil
}

// MaxChar validates that a string must have a length maximum of its constraint
type MaxChar struct {
	Constraint int
	Value      string
}

// Validate check value against constraint
func (v *MaxChar) Validate() error {
	if len(v.Value) > v.Constraint {
		return fmt.Errorf("too long, must be at maximum %v characters", v.Value)
	}

	return nil
}

// Email is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
type Email struct {
	Value string
}

// Validate email addresses
func (v *Email) Validate() error {
	if !strings.Contains(v.Value, "@") || string(v.Value[0]) == "@" || string(v.Value[len(v.Value)-1]) == "@" {
		return fmt.Errorf("invalid email address")
	}

	return nil
}

// Regex allow validation usig regular expressions
type Regex struct {
	Constraint string
	Value      string
}

// Validate using regex
func (v *Regex) Validate() error {
	regex, err := regexp.Compile(v.Constraint)
	if err != nil {
		return err
	}

	if !regex.MatchString(v.Value) {
		return fmt.Errorf("%v doesn't match the pattern %v", v.Value, v.Constraint)
	}

	return nil
}
