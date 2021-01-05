package hijri

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func Test_JD_timeToJulianDays(t *testing.T) {
	jkt := time.FixedZone("WIB", 7*60*60)
	scenarios := []struct {
		date     time.Time
		expected float64
	}{
		{time.Date(-4712, 1, 1, 12, 0, 0, 0, time.UTC), 0},
		{time.Date(-4712, 1, 2, 0, 0, 0, 0, time.UTC), 0.5},
		{time.Date(-4712, 1, 2, 12, 0, 0, 0, time.UTC), 1},
		{time.Date(1582, 10, 4, 0, 0, 0, 0, time.UTC), 2299159.5},
		{time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC), 2299160.5},
		{time.Date(1945, 8, 17, 0, 0, 0, 0, time.UTC), 2431684.5},
		{time.Date(1974, 9, 27, 0, 0, 0, 0, time.UTC), 2442317.5},
		{time.Date(624, 2, 26, 0, 0, 0, 0, time.UTC), 1949029.5},
		{time.Date(-2961, 1, 1, 19, 47, 4, 0, time.UTC), 639553.32435},
		{time.Date(2009, 6, 12, 12, 0, 0, 0, jkt), 2454994.7083},
	}

	for _, scenario := range scenarios {
		jd, _ := timeToJulianDays(scenario.date)
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

func Test_JS_julianDaysToTime(t *testing.T) {
	jkt := time.FixedZone("WIB", 7*60*60)
	scenarios := []struct {
		jd       float64
		expected time.Time
	}{
		{0, time.Date(-4712, 1, 1, 12, 0, 0, 0, time.UTC)},
		{0.5, time.Date(-4712, 1, 2, 0, 0, 0, 0, time.UTC)},
		{1, time.Date(-4712, 1, 2, 12, 0, 0, 0, time.UTC)},
		{2299159.5, time.Date(1582, 10, 4, 0, 0, 0, 0, time.UTC)},
		{2299160.5, time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC)},
		{2431684.5, time.Date(1945, 8, 17, 0, 0, 0, 0, time.UTC)},
		{2442317.5, time.Date(1974, 9, 27, 0, 0, 0, 0, time.UTC)},
		{1949029.5, time.Date(624, 2, 26, 0, 0, 0, 0, time.UTC)},
		{639553.32435, time.Date(-2961, 1, 1, 19, 47, 4, 0, time.UTC)},
		{2454994.70833, time.Date(2009, 6, 12, 12, 0, 0, 0, jkt)},
	}

	for _, scenario := range scenarios {
		date := julianDaysToTime(scenario.jd)
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
