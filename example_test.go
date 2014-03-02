package check

import "fmt"

type Person struct {
	Name string
}

func Example() {
	p := &Person{
		Name: "invalid*",
	}

	s := Struct{
		"Name": Composite{
			NonEmpty{},
			Regex{`^[a-zA-Z0-9]+$`},
			MinChar{10},
		},
	}

	e := s.Validate(*p)

	if e.HasErrors() {
		err, ok := e.GetErrorsByKey("Name")
		if !ok {
			panic("key 'Name' does not exists")
		}
		fmt.Println(err)
	}
}
