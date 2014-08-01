package check

import (
	"strconv"
	"reflect"
)

type Slice struct {
	Inner Validator
}

func (c Slice) Validate(v interface{}) Error {
	err := ValidationError{
		errorMap: map[string][]interface{}{},
	}

	t := reflect.ValueOf(v)

	switch t.Type().Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < t.Len(); i++ {
			if verr := c.Inner.Validate(t.Index(i).Interface()); verr != nil {
				err.errorMap[strconv.Itoa(i)] = append([]interface{}{}, verr)
			}
		}
	default:
		return NewValidationError("notSlice")
	}

	if len(err.errorMap) == 0 {
		return nil
	}

	return err
}
