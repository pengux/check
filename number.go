package check

import "strconv"

// LowerThan validates that a number must be lower than its value
type LowerThan struct {
	Constraint float64
	Value      interface{}
}

// Validate check value against constraint
func (v LowerThan) Validate() Error {
	switch val := v.Value.(type) {
	default:
		return &ValidationError{"NaN", nil}
	case int:
		if v.Constraint <= float64(val) {
			return &ValidationError{"lowerThan", []interface{}{strconv.Itoa(val), strconv.FormatFloat(v.Constraint, 'f', -1, 64)}}
		}
	case float64:
		if v.Constraint <= val {
			return &ValidationError{"lowerThan", []interface{}{strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(v.Constraint, 'f', -1, 64)}}
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
func (v GreaterThan) Validate() Error {
	switch val := v.Value.(type) {
	default:
		return &ValidationError{"NaN", nil}
	case int:
		if v.Constraint >= float64(val) {
			return &ValidationError{"greaterThan", []interface{}{strconv.Itoa(val), strconv.FormatFloat(v.Constraint, 'f', -1, 64)}}
		}
	case float64:
		if v.Constraint >= val {
			return &ValidationError{"greaterThan", []interface{}{strconv.FormatFloat(val, 'f', -1, 64), strconv.FormatFloat(v.Constraint, 'f', -1, 64)}}
		}
	}

	return nil
}
