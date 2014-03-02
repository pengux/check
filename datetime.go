package check

import "time"

// Before check if a time in Value is before the time in Constraint
type Before struct {
	Constraint time.Time
}

// Validate check if a time in Value is before the time in Constraint
func (validator Before) Validate(v interface{}) Error {
	if !v.(time.Time).Before(validator.Constraint) {
		return NewValidationError("before", v.(time.Time).String(), validator.Constraint.String())
	}

	return nil
}

// After check if a time in Value is before the time in Constraint
type After struct {
	Constraint time.Time
}

// Validate check if a time in Value is after the time in Constraint
func (validator After) Validate(v interface{}) Error {
	if !v.(time.Time).After(validator.Constraint) {
		return NewValidationError("after", v.(time.Time).String(), validator.Constraint.String())
	}

	return nil
}
