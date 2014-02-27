package check

import "fmt"

func Example_CustomStringContainValidator() {
	username := "invalid*"
	e := &ErrorMap{}
	e.Add("username", CustomStringContainValidator{"admin", username})
	fmt.Println(e.ToMessages(ErrorMessages))
}
