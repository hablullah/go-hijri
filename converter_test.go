package hijri

import (
	"testing"
	"time"
)

func TestToGregorian(t *testing.T) {
	scenarios := []struct {
		year     int
		month    int
		day      int
		expected time.Time
	}{{
		year:     615,
		month:    9,
		day:      17,
		expected: time.Date(1218, 12, 7, 0, 0, 0, 0, time.UTC),
	}, {
		year:     1430,
		month:    1,
		day:      1,
		expected: time.Date(2008, 12, 29, 0, 0, 0, 0, time.UTC),
	}, {
		year:     -1000,
		month:    1,
		day:      1,
		expected: time.Date(-349, 5, 14, 0, 0, 0, 0, time.UTC),
	}, {
		year:     -640,
		month:    5,
		day:      16,
		expected: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
	}, {
		year:     100,
		month:    10,
		day:      1,
		expected: time.Date(719, 4, 26, 0, 0, 0, 0, time.UTC),
	}, {
		year:     990,
		month:    9,
		day:      17,
		expected: time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC),
	}, {
		year:     1502,
		month:    12,
		day:      30,
		expected: time.Date(2079, 10, 26, 0, 0, 0, 0, time.UTC),
	}}

	for _, scenario := range scenarios {
		date := ToGregorian(scenario.year, scenario.month, scenario.day)
		diff := date.Sub(scenario.expected).Hours()

		if diff != 0 {
			t.Errorf("\n"+
				"hijri    : %04d-%02d-%02d\n"+
				"expected : %s\n"+
				"get      : %s",
				scenario.year, scenario.month, scenario.day,
				scenario.expected.Format("2006-01-02"),
				date.Format("2006-01-02"))
		}
	}
}

func TestToHijri(t *testing.T) {
	scenarios := []struct {
		date   time.Time
		hYear  int
		hMonth int
		hDay   int
	}{{
		date:   time.Date(1218, 12, 7, 0, 0, 0, 0, time.UTC),
		hYear:  615,
		hMonth: 9,
		hDay:   17,
	}, {
		date:   time.Date(2008, 12, 29, 0, 0, 0, 0, time.UTC),
		hYear:  1430,
		hMonth: 1,
		hDay:   1,
	}, {
		date:   time.Date(-349, 5, 14, 0, 0, 0, 0, time.UTC),
		hYear:  -1000,
		hMonth: 1,
		hDay:   1,
	}, {
		date:   time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		hYear:  -640,
		hMonth: 5,
		hDay:   16,
	}, {
		date:   time.Date(719, 4, 26, 0, 0, 0, 0, time.UTC),
		hYear:  100,
		hMonth: 10,
		hDay:   1,
	}, {
		date:   time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC),
		hYear:  990,
		hMonth: 9,
		hDay:   17,
	}, {
		date:   time.Date(2079, 10, 26, 0, 0, 0, 0, time.UTC),
		hYear:  1502,
		hMonth: 12,
		hDay:   30,
	}}

	for _, scenario := range scenarios {
		Y, M, D := ToHijri(scenario.date)

		if Y != scenario.hYear || M != scenario.hMonth || D != scenario.hDay {
			t.Errorf("\n"+
				"date     : %s\n"+
				"expected : %04d-%02d-%02d\n"+
				"get      : %04d-%02d-%02d",
				scenario.date.Format("2006-01-02"),
				scenario.hYear, scenario.hMonth, scenario.hDay,
				Y, M, D)
		}
	}
}
