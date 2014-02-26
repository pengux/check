package govalid

import (
	"errors"
	"fmt"
)

// LowerThan validates that a number must be lower than its value
type LowerThan struct {
	Constraint float64
	Value      interface{}
}

// Validate check value against constraint
func (v *LowerThan) Validate() error {
	switch val := v.Value.(type) {
	default:
		return errors.New("value is not a number")
	case int:
		if v.Constraint <= float64(val) {
			return fmt.Errorf("%v is not lower than %v", val, v.Constraint)
		}
	case float64:
		if v.Constraint <= val {
			return fmt.Errorf("%v is not lower than %v", val, v.Constraint)
		}
	}

	return nil
}

// GreaterThan validates that a number must be greater than its value
type GreaterThan struct {
	Constraint float64
	Value      interface{}
}

// Validate check value against constraint
func (v *GreaterThan) Validate() error {
	switch val := v.Value.(type) {
	default:
		return errors.New("value is not a number")
	case int:
		if v.Constraint >= float64(val) {
			return fmt.Errorf("%v is not greater than %v", val, v.Constraint)
		}
	case float64:
		if v.Constraint >= val {
			return fmt.Errorf("%v is not greater than %v", val, v.Constraint)
		}
	}

	return nil
}
