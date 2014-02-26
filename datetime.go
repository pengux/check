package govalid

import (
	"fmt"
	"time"
)

// Before check if a time in Value is before the time in Constraint
type Before struct {
	Constraint time.Time
	Value      time.Time
}

// Validate check if a time in Value is before the time in Constraint
func (v *Before) Validate() error {
	if !v.Value.Before(v.Constraint) {
		return fmt.Errorf("%v is not before %v", v.Value, v.Constraint)
	}

	return nil
}

// After check if a time in Value is before the time in Constraint
type After struct {
	Constraint time.Time
	Value      time.Time
}

// Validate check if a time in Value is after the time in Constraint
func (v *After) Validate() error {
	if !v.Value.After(v.Constraint) {
		return fmt.Errorf("%v is not after %v", v.Value, v.Constraint)
	}

	return nil
}
