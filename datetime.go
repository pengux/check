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
func (v Before) Validate() (err error, params []string) {
	if !v.Value.Before(v.Constraint) {
		return fmt.Errorf("before"), []string{v.Value.String(), v.Constraint.String()}
	}

	return nil, params
}

// After check if a time in Value is before the time in Constraint
type After struct {
	Constraint time.Time
	Value      time.Time
}

// Validate check if a time in Value is after the time in Constraint
func (v After) Validate() (err error, params []string) {
	if !v.Value.After(v.Constraint) {
		return fmt.Errorf("after"), []string{v.Value.String(), v.Constraint.String()}
	}

	return nil, params
}
