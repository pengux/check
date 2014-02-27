package check

import (
	"testing"
	"time"
)

func TestDatetimeValidators(t *testing.T) {
	var validatorTests = []testTable{
		{&Before{
			time.Date(2014, time.January, 1, 1, 0, 0, 0, time.UTC),
			time.Date(2013, time.January, 1, 1, 0, 0, 0, time.UTC),
		}, true, "Expected '2013-01-01 01:00:00' to be before '2014-01-01 01:00:00:00'"},
		{&After{
			time.Date(2014, time.January, 1, 1, 0, 0, 0, time.UTC),
			time.Date(2015, time.January, 1, 1, 0, 0, 0, time.UTC),
		}, true, "Expected '2015-01-01 01:00:00' to be after '2014-01-01 01:00:00:00'"},
	}

	runTableTest(validatorTests, t)
}
