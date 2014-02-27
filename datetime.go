package check

import "time"

// Before check if a time in Value is before the time in Constraint
type Before struct {
	Constraint time.Time
	Value      time.Time
}

// Validate check if a time in Value is before the time in Constraint
func (v Before) Validate() Error {
	if !v.Value.Before(v.Constraint) {
		return &ValidationError{"before", []interface{}{v.Value.String(), v.Constraint.String()}}
	}

	return nil
}

// After check if a time in Value is before the time in Constraint
type After struct {
	Constraint time.Time
	Value      time.Time
}

// Validate check if a time in Value is after the time in Constraint
func (v After) Validate() Error {
	if !v.Value.After(v.Constraint) {
		return &ValidationError{"after", []interface{}{v.Value.String(), v.Constraint.String()}}
	}

	return nil
}
