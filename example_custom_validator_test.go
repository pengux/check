package check

import "fmt"

func ExampleValidator() {
	username := "invalid*"
	validator := CustomStringContainValidator{"admin"}
	e := validator.Validate(username)
	fmt.Println(ErrorMessages[e.Error()])
}
