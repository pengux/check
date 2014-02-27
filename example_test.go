package check

import "fmt"

func Example() {
	username := "invalid*"
	e := &ErrorMap{}
	e.Add("username", Regex{"[a-zA-Z0-9]+$", username})
	e.Add("username", NonEmpty{username}, MinChar{10, username}) // Add multiple validators at the same time

	if e.HasErrors() {
		err, ok := e.GetErrorsByKey("username")
		if !ok {
			panic("key username does not exists")
		}
		fmt.Println(err)
	}
}
