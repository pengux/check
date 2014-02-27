package govalid

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

type CustomStringContainValidator struct {
	Constraint string
	Value      string
}

func (v *CustomStringContainValidator) Validate() (err error, params []string) {
	if !strings.Contains(v.Value, v.Constraint) {
		return fmt.Errorf("customStringContainValidator"), []string{v.Value, v.Constraint}
	}

	return nil, params
}

type User struct {
	Username string
	Password string
	Name     string
	Age      int
	Email    string
	Birthday time.Time
}

func (u *User) Validate() (ValidationErrors, bool) {
	return Validate(map[string][]Validator{
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
				8,
				u.Password,
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
	})
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

	errs, hasErr := invalidUser.Validate()
	if !hasErr {
		t.Errorf("Expected 'invalidUser' to be invalid")
	}

	errMessages := errs.ToMessages(ErrorMessages)
	if errMessages["name"]["nonZero"] != ErrorMessages["nonZero"] {
		t.Errorf("Expected proper error message")
	}

	json, _ := json.MarshalIndent(errMessages, "", "	")
	log.Println(string(json))

	_, hasErr = validUser.Validate()
	if hasErr {
		t.Errorf("Expected 'validUser' to be valid")
	}
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		invalidUser := &User{
			"not-valid-username*",
			"123",   // Invalid password length
			"",      // Cannot be empty
			150,     // Invalid age
			"@test", // Invalid email address
			time.Date(1991, time.January, 1, 1, 0, 0, 0, time.UTC), // Invalid date
		}

		invalidUser.Validate()
	}
}
