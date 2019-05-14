package hijri

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func Test_dateToJD(t *testing.T) {
	jkt := time.FixedZone("WIB", 7*60*60)
	scenarios := []struct {
		date     time.Time
		expected float64
	}{{
		date:     time.Date(-4712, 1, 1, 12, 0, 0, 0, time.UTC),
		expected: 0,
	}, {
		date:     time.Date(-4712, 1, 2, 0, 0, 0, 0, time.UTC),
		expected: 0.5,
	}, {
		date:     time.Date(-4712, 1, 2, 12, 0, 0, 0, time.UTC),
		expected: 1,
	}, {
		date:     time.Date(1582, 10, 4, 0, 0, 0, 0, time.UTC),
		expected: 2299159.5,
	}, {
		date:     time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC),
		expected: 2299160.5,
	}, {
		date:     time.Date(1945, 8, 17, 0, 0, 0, 0, time.UTC),
		expected: 2431684.5,
	}, {
		date:     time.Date(1974, 9, 27, 0, 0, 0, 0, time.UTC),
		expected: 2442317.5,
	}, {
		date:     time.Date(624, 2, 26, 0, 0, 0, 0, time.UTC),
		expected: 1949029.5,
	}, {
		date:     time.Date(-2961, 1, 1, 19, 47, 4, 0, time.UTC),
		expected: 639553.32435,
	}, {
		date:     time.Date(2009, 6, 12, 12, 0, 0, 0, jkt),
		expected: 2454994.7083,
	}}

	for _, scenario := range scenarios {
		jd := dateToJD(scenario.date)
		diff := decimal.NewFromFloat(jd).
			Sub(decimal.NewFromFloat(scenario.expected))

		if !diff.Round(3).Equal(decimal.Zero) {
			t.Errorf("\n"+
				"date     : %s\n"+
				"expected : %f\n"+
				"get      : %f",
				scenario.date.Format("2006-01-02 15:04:05 -07"),
				scenario.expected,
				jd)
		}
	}
}

func Test_jdToDate(t *testing.T) {
	jkt := time.FixedZone("WIB", 7*60*60)
	scenarios := []struct {
		jd       float64
		expected time.Time
	}{{
		jd:       0,
		expected: time.Date(-4712, 1, 1, 12, 0, 0, 0, time.UTC),
	}, {
		jd:       0.5,
		expected: time.Date(-4712, 1, 2, 0, 0, 0, 0, time.UTC),
	}, {
		jd:       1,
		expected: time.Date(-4712, 1, 2, 12, 0, 0, 0, time.UTC),
	}, {
		jd:       2299159.5,
		expected: time.Date(1582, 10, 4, 0, 0, 0, 0, time.UTC),
	}, {
		jd:       2299160.5,
		expected: time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC),
	}, {
		jd:       2431684.5,
		expected: time.Date(1945, 8, 17, 0, 0, 0, 0, time.UTC),
	}, {
		jd:       2442317.5,
		expected: time.Date(1974, 9, 27, 0, 0, 0, 0, time.UTC),
	}, {
		jd:       1949029.5,
		expected: time.Date(624, 2, 26, 0, 0, 0, 0, time.UTC),
	}, {
		jd:       639553.32435,
		expected: time.Date(-2961, 1, 1, 19, 47, 4, 0, time.UTC),
	}, {
		jd:       2454994.70833,
		expected: time.Date(2009, 6, 12, 12, 0, 0, 0, jkt),
	}}

	for _, scenario := range scenarios {
		date := jdToDate(scenario.jd)
		diff := date.Sub(scenario.expected).Seconds()

		// Since float is fickle, allow +-1 second tolerance
		if diff < -1 || diff > 1 {
			t.Errorf("\n"+
				"JD       : %.03f\n"+
				"expected : %s\n"+
				"get      : %s\n",
				scenario.jd,
				scenario.expected.UTC().Format("2006-01-02 15:04:05"),
				date.Format("2006-01-02 15:04:05"))
		}
	}
}
