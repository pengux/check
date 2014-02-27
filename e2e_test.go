package check

import (
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

func (v CustomStringContainValidator) Validate() (err error, params []string) {
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

func (u *User) Validate() *ErrorMap {
	e := &ErrorMap{}
	e.Add("username", NonZero{u.Username})
	e.Add("username", Regex{`^[a-zA-Z0-9]+$`, u.Username})
	e.Add("password", NonZero{u.Password}, MinChar{8, u.Password})
	e.Add("name", NonZero{u.Name})
	e.Add("age", GreaterThan{3, u.Age}, LowerThan{120, u.Age})
	e.Add("email", Email{u.Email})
	e.Add("email", CustomStringContainValidator{"test.com", u.Email})
	e.Add("birthday", Before{time.Date(1990, time.January, 1, 1, 0, 0, 0, time.UTC), u.Birthday}, After{time.Date(1900, time.January, 1, 1, 0, 0, 0, time.UTC), u.Birthday})

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

	err, ok := e.GetErrorsByKey("username")
	if !ok {
		t.Errorf("Expected errors for 'username'")
	} else {
		if len(err) < 1 {
			t.Errorf("Expected 1 error for 'username'")
		}
	}

	errMessages := e.ToMessages(ErrorMessages)
	if errMessages["name"]["nonZero"] != ErrorMessages["nonZero"] {
		t.Errorf("Expected proper error message")
	}

	// json, _ := json.MarshalIndent(errMessages, "", "	")
	// log.Println(string(json))

	e = validUser.Validate()
	if e.HasErrors() {
		log.Println(*e)
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
