package govalid

import (
	"encoding/json"
	"log"
	"testing"
)

type User struct {
	Username string
	Password string
	Name     string
	Age      int
	Email    string
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
	}

	validUser := &User{
		"testuser",
		"validPassword123",
		"Good Name",
		20,
		"test@test.com",
	}

	errs, hasErr := Validate(invalidUser)
	if !hasErr {
		t.Errorf("Expected 'invalidUser' to be invalid")
	}

	log.Println(errs)
	json, _ := json.MarshalIndent(errs, "", "	")
	log.Println(string(json))

	errs, hasErr = Validate(validUser)
	if hasErr {
		t.Errorf("Expected 'validUser' to be valid")
	}
}
