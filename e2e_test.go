package check

import (
	"strings"
	"testing"
	"time"
)

type CustomStringContainValidator struct {
	Constraint string
}

func (validator CustomStringContainValidator) Validate(v interface{}) Error {
	if !strings.Contains(v.(string), validator.Constraint) {
		return &ValidationError{
			map[string][]interface{}{
				"ecustomStringContainValidato": []interface{}{
					v.(string),
					validator.Constraint,
				},
			},
		}
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

func (u *User) Validate() StructError {
	s := Struct{
		"Username": Composite{
			NonEmpty{},
			Regex{`^[a-zA-Z0-9]+$`},
		},
		"Password": Composite{
			NonEmpty{},
			MinChar{8},
		},
		"Name": NonEmpty{},
		"Age": Composite{
			GreaterThan{3},
			LowerThan{120},
		},
		"Email": Composite{
			Email{},
			CustomStringContainValidator{"test.com"},
		},
		"Birthday": Composite{
			Before{time.Date(1990, time.January, 1, 1, 0, 0, 0, time.UTC)},
			After{time.Date(1900, time.January, 1, 1, 0, 0, 0, time.UTC)},
		},
	}
	e := s.Validate(*u)

	return e
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

	e := invalidUser.Validate()
	if !e.HasErrors() {
		t.Errorf("Expected 'invalidUser' to be invalid")
	}

	err, ok := e.GetErrorsByKey("Username")
	if !ok {
		t.Errorf("Expected errors for 'Username'")
	} else {
		if len(err) < 1 {
			t.Errorf("Expected 1 error for 'Username'")
		}
	}

	errMessages := e.ToMessages(ErrorMessages)
	if errMessages["Name"]["nonZero"] != ErrorMessages["nonZero"] {
		t.Errorf("Expected proper error message")
	}

	// json, _ := json.MarshalIndent(errMessages, "", "	")
	// log.Println(string(json))

	e = validUser.Validate()
	if e.HasErrors() {
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
