package hijri

import (
	"fmt"
	"testing"
	"time"
)

func TestToGregorian(t *testing.T) {
	scenarios := []struct {
		arg      string
		expected time.Time
	}{{
		arg:      "615-9-17",
		expected: time.Date(1218, 12, 7, 0, 0, 0, 0, time.UTC),
	}, {
		arg:      "1430-1-1",
		expected: time.Date(2008, 12, 29, 0, 0, 0, 0, time.UTC),
	}, {
		arg:      "-1000-1-1",
		expected: time.Date(-349, 5, 14, 0, 0, 0, 0, time.UTC),
	}, {
		arg:      "-640-5-16",
		expected: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
	}, {
		arg:      "100-10-1",
		expected: time.Date(719, 4, 26, 0, 0, 0, 0, time.UTC),
	}, {
		arg:      "990-9-17",
		expected: time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC),
	}, {
		arg:      "1502-12-30",
		expected: time.Date(2079, 10, 26, 0, 0, 0, 0, time.UTC),
	}}

	for _, scenario := range scenarios {
		var hYear, hMonth, hDay int
		_, err := fmt.Sscanf(scenario.arg, "%d-%d-%d", &hYear, &hMonth, &hDay)
		if err != nil {
			t.Errorf("failed to parse arg: %v", err)
		}

		date := ToGregorian(hYear, hMonth, hDay)
		diff := date.Sub(scenario.expected).Hours()

		if diff != 0 {
			t.Errorf("\n"+
				"hijri    : %s\n"+
				"expected : %s\n"+
				"get      : %s",
				scenario.arg,
				scenario.expected.Format("2006-01-02"),
				date.Format("2006-01-02"))
		}
	}
}

func TestToHijri(t *testing.T) {
	scenarios := []struct {
		arg      time.Time
		expected string
	}{{
		arg:      time.Date(1218, 12, 7, 0, 0, 0, 0, time.UTC),
		expected: "615-9-17",
	}, {
		arg:      time.Date(2008, 12, 29, 0, 0, 0, 0, time.UTC),
		expected: "1430-1-1",
	}, {
		arg:      time.Date(-349, 5, 14, 0, 0, 0, 0, time.UTC),
		expected: "-1000-1-1",
	}, {
		arg:      time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		expected: "-640-5-16",
	}, {
		arg:      time.Date(719, 4, 26, 0, 0, 0, 0, time.UTC),
		expected: "100-10-1",
	}, {
		arg:      time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC),
		expected: "990-9-17",
	}, {
		arg:      time.Date(2079, 10, 26, 0, 0, 0, 0, time.UTC),
		expected: "1502-12-30",
	}}

	for _, scenario := range scenarios {
		year, month, day := ToHijri(scenario.arg)
		result := fmt.Sprintf("%d-%d-%d", year, month, day)

		if result != scenario.expected {
			t.Errorf("\n"+
				"date     : %s\n"+
				"expected : %s\n"+
				"get      : %s",
				scenario.arg.Format("2006-01-02"),
				scenario.expected, result)
		}
	}
}

func TestToUmmAlQura(t *testing.T) {
	scenarios := []struct {
		arg      time.Time
		expected string
	}{{
		arg:      time.Date(1937, 3, 14, 0, 0, 0, 0, time.UTC),
		expected: "1356-1-1",
	}, {
		arg:      time.Date(2077, 11, 16, 0, 0, 0, 0, time.UTC),
		expected: "1500-12-30",
	}, {
		arg:      time.Date(2008, 12, 29, 0, 0, 0, 0, time.UTC),
		expected: "1430-1-1",
	}, {
		arg:      time.Date(2019, 9, 12, 17, 0, 0, 0, time.Local),
		expected: "1441-1-13",
	}}

	for _, scenario := range scenarios {
		year, month, day, _ := ToUmmAlQura(scenario.arg)
		result := fmt.Sprintf("%d-%d-%d", year, month, day)

		if result != scenario.expected {
			t.Errorf("\n"+
				"date     : %s\n"+
				"expected : %s\n"+
				"get      : %s",
				scenario.arg.Format("2006-01-02"),
				scenario.expected, result)
		}
	}
}
