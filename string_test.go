package check

import "testing"

func TestStringValidators(t *testing.T) {
	var validatorTests = []testTable{
		{&MinChar{3}, "foo", true, "Expected 'foo' to be minimum of 3 char"},
		{&MinChar{5}, "foo", false, "Expected 'foo' NOT to be minimum of 5 char"},
		{&MaxChar{3}, "foo", true, "Expected 'foo' to be maximum of 3 char"},
		{&MaxChar{2}, "foo", false, "Expected 'foo' NOT to be maximum of 2 char"},
		{&Email{}, "@foo", false, "Expected '@test' to be invalid email addressy"},
		{&Email{}, "foo@", false, "Expected 'test@' to be invalid email addressy"},
		{&Email{}, "foo", false, "Expected 'test' to be invalid email addressy"},
		{&Email{}, "foo@bar", true, "Expected 'foo@bar' to be a valid email address"},
		{&Regex{"[a-zA-Z0-9]"}, "aA0", true, "Expected 'aA0' to match the regex '[a-zA-Z0-9]'"},
		{&Regex{"[a-zA-Z0-9]"}, "*", false, "Expected '*' NOT to match the regex '[a-zA-Z0-9]'"},
	}

	runTableTest(validatorTests, t)
}
