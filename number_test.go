package check

import "testing"

func TestNumberValidators(t *testing.T) {
	var validatorTests = []testTable{
		{&LowerThan{2, 1}, true, "Expected 1 to be lower than 2"},
		{&LowerThan{1, 2}, false, "Expected 2 NOT to be lower than 1"},
		{&LowerThan{2, 2}, false, "Expected 2 NOT to be lower than 2"},
		{&GreaterThan{1, 2}, true, "Expected 2 to be greater than 1"},
		{&GreaterThan{2, 1}, false, "Expected 1 NOT to be greater than 2"},
		{&GreaterThan{2, 2}, false, "Expected 2 NOT to be greater than 2"},
	}

	runTableTest(validatorTests, t)
}
