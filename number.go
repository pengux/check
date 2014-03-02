package check

import "strconv"

// LowerThan validates that a number must be lower than its value
type LowerThan struct {
	Constraint float64
}

// Validate check value against constraint
func (validator LowerThan) Validate(v interface{}) Error {
	switch val := v.(type) {
	default:
		return &ValidationError{map[string][]interface{}{"NaN": nil}}
	case int:
		if validator.Constraint <= float64(val) {
			return &ValidationError{
				map[string][]interface{}{
					"lowerThan": []interface{}{
						strconv.Itoa(val),
						strconv.FormatFloat(validator.Constraint, 'f', -1, 64),
					},
				},
			}
		}
	case float64:
		if validator.Constraint <= val {
			return &ValidationError{
				map[string][]interface{}{
					"lowerThan": []interface{}{
						strconv.FormatFloat(val, 'f', -1, 64),
						strconv.FormatFloat(validator.Constraint, 'f', -1, 64),
					},
				},
			}
		}
	}

	return nil
}

// GreaterThan validates that a number must be greater than its value
type GreaterThan struct {
	Constraint float64
}

// Validate check value against constraint
func (validator GreaterThan) Validate(v interface{}) Error {
	switch val := v.(type) {
	default:
		return &ValidationError{map[string][]interface{}{"NaN": nil}}
	case int:
		if validator.Constraint >= float64(val) {
			return &ValidationError{
				map[string][]interface{}{
					"greaterThan": []interface{}{
						strconv.Itoa(val),
						strconv.FormatFloat(validator.Constraint, 'f', -1, 64),
					},
				},
			}
		}
	case float64:
		if validator.Constraint >= val {
			return &ValidationError{
				map[string][]interface{}{
					"greaterThan": []interface{}{
						strconv.FormatFloat(val, 'f', -1, 64),
						strconv.FormatFloat(validator.Constraint, 'f', -1, 64),
					},
				},
			}
		}
	}

	return nil
}
