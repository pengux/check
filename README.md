
# check - Go package for data validation
[Documentation](http://godoc.org/github.com/pengux/check)

[![Build Status](https://travis-ci.org/pengux/check.svg?branch=master)](https://travis-ci.org/pengux/check) [![GoDoc](https://godoc.org/github.com/pengux/check?status.svg)](https://godoc.org/github.com/pengux/check)

## Design goals
- Composite pattern
- Multiple constraints on the same value by applying multiple validators
- Easy to create custom validators
- Easy to customize error messages

## Usage
```bash
go get github.com/pengux/check
```


To run tests:
```bash
cd $GOPATH/src/github.com/pengux/check && go test
```


To validate your data, create a new Struct and add validators to it:

```go
type User struct {
	Username string
}

func main() {
	u := &User{
		Username: "invalid*",
	}

	s := check.Struct{
		"Username": check.Composite{
			check.NonEmpty{},
			check.Regex{`^[a-zA-Z0-9]+$`},
			check.MinChar{10},
		},
	}

	e := s.Validate(u)

	if e.HasErrors() {
		err, ok := e.GetErrorsByKey("Username")
		if !ok {
			panic("key 'Username' does not exists")
		}
		fmt.Println(err)
	}
}
```

To use your own custom validator, just implement the Validator interface:

```go
type CustomStringContainValidator struct {
	Constraint string
}

func (validator CustomStringContainValidator) Validate(v interface{}) check.Error {
	if !strings.Contains(v.(string), validator.Constraint) {
		return check.NewValidationError("customStringContainValidator", v, validator.Constraint)
	}

	return nil
}

func main() {
	username := "invalid*"
	validator := CustomStringContainValidator{"admin"}
	e := validator.Validate(username)
	fmt.Println(check.ErrorMessages[e.Error()])
}
```

To use custom error messages, either overwrite the package variable `ErrorMessages` or create your own `map[string]string`:

```go
check.ErrorMessages["minChar"] := "the string must be minimum %v characters long"
errMessages := errs.ToMessages()
fmt.Println(errMessages)
```

For more example code check the file [`e2e_test.go`](https://github.com/pengux/check/blob/master/e2e_test.go).

