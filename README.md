[![baby-gopher](https://raw2.github.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

# check - Go package for data validation

## Design goals
- Multiple constraints on the same value by applying multiple validators
- Easy to create custom validators
- Easy to customize error messages
- Minimal usage of reflect package

## Usage
```bash
go get github.com/pengux/check
```


To run tests:
```bash
cd $GOPATH/src/github.com/pengux/check && go test
```


To validate your data, create a new ErrorMap and add validators to it:

```go
func main() {
	username := "invalid*"
	e := &check.ErrorMap{}
	e.Add("username", check.Regex{"[a-zA-Z0-9]+$", username})
	e.Add("username", check.NonEmpty{username}, check.MinChar{10, username}) // Add multiple validators at the same time

	if e.HasErrors() {
		err, ok := e.GetErrorsByKey("username")
		if !ok {
			panic("key username does not exists")
		}
		fmt.Println(err)
	}
}
```


To use your own custom validator, just implement the Validator interface:

```go
type CustomStringContainValidator struct {
	Constraint string
	Value      string
}

func (v CustomStringContainValidator) Validate() check.Error {
	if !strings.Contains(v.Value, v.Constraint) {
		return &check.ValidationError{"customStringContainValidator", []interface{}{v.Value, v.Constraint}}
	}

	return nil
}

func main() {
	username := "invalid*"
	e := &check.ErrorMap{}
	e.Add("username", CustomStringContainValidator{"admin", username})
	fmt.Println(e.ToMessages(check.ErrorMessages))
}
```


To use default error messages, pass in the package variable ErrorMessages:

```go
errMessages := e.ToMessages(check.ErrorMessages)
fmt.Println(errMessages)
```


To use custom error messages, either overwrite the package variable `ErrorMessages` or create your own `map[string]string`:

```go
check.ErrorMessages["minChar"] := "the string must be minimum %v characters long"
errMessages := errs.ToMessages(check.ErrorMessages)
fmt.Println(errMessages)

errMessages := errs.ToMessages(map[string]string{"minChar": "the string must be minimum %v characters long"})
fmt.Println(errMessages)
```


For more example code check the file [`e2e_test.go`](https://github.com/pengux/check/blob/master/e2e_test.go).

