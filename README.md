[![baby-gopher](https://raw2.github.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

# govalid - Go package for data validation

## Design goals
- Multiple constraints on the same value by applying multiple validators
- Easy to create custom validators
- Easy to customize error messages
- Minimal usage of reflect package

## Usage
```bash
go get github.com/pengux/govalid
```


To run tests:
```bash
cd $GOPATH/src/github.com/pengux/govalid && go test
```


To validate your data, you need to create a map with fields and validators and pass it to the Validate() function:

```go
func main() {
	errs, hasErr := govalid.Validate(map[string][]Validator{
		"foo": []govalid.Validator{
			govalid.Regex{
				`[a-zA-Z0-9]`, // constraint
				"invalid-string", // value to be validated
			},
			govalid.NonZero{
				"", // Invalid
			},
			govalid.MinChar{
				5,
				"bar", // Invalid
			},
		}
	})
	fmt.Println(errs, hasErr)
}
```


To use your own custom validator, just implement the Validator interface:

```go
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

func main() {
	errs, hasErr := govalid.Validate(map[string][]Validator{
		"foo": []govalid.Validator{
			CustomStringContainValidator{
				"test.com",
				"foo@bar.com",
			},
		}
	})
	fmt.Println(errs, hasErr)
}
```


To use default error messages, pass in the package variable ErrorMessages:

```go
errMessages := errs.ToMessages(govalid.ErrorMessages)
fmt.Println(errMessages)
```


To use custom error messages, either overwrite the package variable `ErrorMessages` or create your `map[string]string`:

```go
govalid.ErrorMessages["minChar"] := "the string must be minimum %v characters long"
errMessages := errs.ToMessages(govalid.ErrorMessages)
fmt.Println(errMessages)

errMessages := errs.ToMessages(map[string]string{"minChar": "the string must be minimum %v characters long"})
fmt.Println(errMessages)
```


For more example code check the file [`e2e_test.go`](https://github.com/pengux/govalid/blob/master/e2e_test.go).

