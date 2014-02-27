package govalid

import (
	"fmt"
	"strconv"
)

// LowerThan validates that a number must be lower than its value
type LowerThan struct {
	Constraint float64
	Value      interface{}
}

// Validate check value against constraint
func (v *LowerThan) Validate() (err error, params []string) {
	switch val := v.Value.(type) {
	default:
		return fmt.Errorf("nan"), params
	case int:
		if v.Constraint <= float64(val) {
			return fmt.Errorf("lowerThan"), []string{strconv.Itoa(val), strconv.FormatFloat(v.Constraint, 'f', -1, 64)}
		}
	case float64:
		if v.Constraint <= val {
			return fmt.Errorf("lowerThan"), []string{strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(v.Constraint, 'f', -1, 64)}
		}
	}

	return nil, params
}

// GreaterThan validates that a number must be greater than its value
type GreaterThan struct {
	Constraint float64
	Value      interface{}
}

// Validate check value against constraint
func (v *GreaterThan) Validate() (err error, params []string) {
	switch val := v.Value.(type) {
	default:
		return fmt.Errorf("nan"), params
	case int:
		if v.Constraint >= float64(val) {
			return fmt.Errorf("greaterThan"), []string{strconv.Itoa(val), strconv.FormatFloat(v.Constraint, 'f', -1, 64)}
		}
	case float64:
		if v.Constraint >= val {
			return fmt.Errorf("greaterThan"), []string{strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(v.Constraint, 'f', -1, 64)}
		}
	}

	return nil, params
}
