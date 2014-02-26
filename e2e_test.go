package govalid

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type CustomStringContainValidator struct {
	Constraint string
	Value      string
}

func (v *CustomStringContainValidator) Validate() error {
	if !strings.Contains(v.Value, v.Constraint) {
		return fmt.Errorf("the string %v doesn't contain %v", v.Value, v.Constraint)
	}

	return nil
}

type User struct {
	Username string
	Password string
	Name     string
	Age      int
	Email    string
	Birthday time.Time
}

func (u *User) GetValidators() map[string][]Validator {
	return map[string][]Validator{
		"username": []Validator{
			&NonZero{
				u.Username,
			},
			&Regex{
				`[a-zA-Z0-9]`, // constraint
				u.Username,    // value to be validated
			},
		},
		"password": []Validator{
			&NonZero{
				u.Password,
			},
			&MinChar{
				8,          // constraint
				u.Password, // value to be validated
			},
		},
		"name": []Validator{
			&NonZero{
				u.Name,
			},
		},
		"age": []Validator{
			&GreaterThan{
				3,
				u.Age,
			},
			&LowerThan{
				120,
				u.Age,
			},
		},
		"email": []Validator{
			&Email{
				u.Email,
			},
			&CustomStringContainValidator{
				"test.com",
				u.Email,
			},
		},
		"birthday": []Validator{
			&Before{
				time.Date(1990, time.January, 1, 1, 0, 0, 0, time.UTC),
				u.Birthday,
			},
			&After{
				time.Date(1900, time.January, 1, 1, 0, 0, 0, time.UTC),
				u.Birthday,
			},
		},
	}
}

func TestIntegration(t *testing.T) {
	invalidUser := &User{
		"not-valid-username*",
		"123",   // Invalid password length
		"",      // Cannot be empty
		150,     // Invalid age
		"@test", // Invalid email address
		time.Date(1991, time.January, 1, 1, 0, 0, 0, time.UTC), // Invalid date
	}

	validUser := &User{
		"testuser",
		"validPassword123",
		"Good Name",
		20,
		"test@test.com",
		time.Date(1980, time.January, 1, 1, 0, 0, 0, time.UTC),
	}

	_, hasErr := Validate(invalidUser)
	if !hasErr {
		t.Errorf("Expected 'invalidUser' to be invalid")
	}

	// errs, hasErr := Validate(invalidUser)
	// Marshal errors into json
	// json, _ := json.MarshalIndent(errs, "", "	")
	// log.Println(string(json))

	_, hasErr = Validate(validUser)
	if hasErr {
		t.Errorf("Expected 'validUser' to be valid")
	}
}
