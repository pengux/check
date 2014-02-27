package check

import "testing"

type foo struct {
	b bool
}

type testTable struct {
	v        Validator
	valid    bool
	expected string
}

func runTableTest(table []testTable, t *testing.T) {
	for _, tt := range table {
		err := tt.v.Validate()
		if (err != nil && tt.valid) || (err == nil && !tt.valid) {
			t.Errorf(tt.expected)
		}
	}
}

func TestNonZero(t *testing.T) {
	var validatorTests = []testTable{
		{&NonZero{int(1)}, true, "Expected int 1 to be non-zero"},
		{&NonZero{float64(1.0)}, true, "Expected float64 1.0 to be non-zero"},
		{&NonZero{"foo"}, true, "Expected string 'foo' to be non-zero"},
		{&NonZero{true}, true, "Expected boolean true to be non-zero"},
		{&NonZero{foo{true}}, true, "Expected struct 'foo' with value to be non-zero"},
		{&NonZero{[]foo{foo{true}}}, true, "Expected slice of structs 'foo' with value to be non-zero"},
		{&NonZero{int(0)}, false, "Expected int 0 to be zero"},
		{&NonZero{float64(0.0)}, false, "Expected float64 0.0 to be zero"},
		{&NonZero{""}, false, "Expected string '' to be zero"},
		{&NonZero{false}, false, "Expected boolean false to be zero"},
		{&NonZero{foo{}}, false, "Expected struct 'foo' with no value to be zero"},
		{&NonZero{[]foo{}}, false, "Expected slice of structs 'foo' with no value to be zero"},
	}

	runTableTest(validatorTests, t)
}
