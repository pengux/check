package check

import "fmt"

func ExampleValidator() {
	username := "invalid*"
	e := &ErrorMap{}
	e.Add("username", CustomStringContainValidator{"admin", username})
	fmt.Println(e.ToMessages(ErrorMessages))
}
