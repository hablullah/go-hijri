package hijri_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RadhiFadlillah/go-hijri"
)

var hijriTestData []TestData

func init() {
	var err error
	hijriTestData, err = generateTestData("test/hijri.csv")
	if err != nil {
		panic(err)
	}
}

func Test_Hijri_ConvertDate(t *testing.T) {
	if len(hijriTestData) == 0 {
		t.Fatal("no tests available for Hijri")
	}

	for _, data := range hijriTestData {
		gregorianDate, _ := time.Parse("2006-01-02", data.Gregorian)
		hijriDate, _ := hijri.CreateHijriDate(gregorianDate, hijri.Default)
		strHijriDate := fmt.Sprintf("%04d-%02d-%02d",
			hijriDate.Year,
			hijriDate.Month,
			hijriDate.Day)

		if strHijriDate != data.Hijri {
			t.Errorf("%s: want %s got %s\n", data.Gregorian, data.Hijri, strHijriDate)
		}
	}
}

func Test_Hijri_ToGregorian(t *testing.T) {
	if len(hijriTestData) == 0 {
		t.Fatal("no tests available for Hijri")
	}

	for _, data := range hijriTestData {
		var hijriDate hijri.HijriDate
		fmt.Sscanf(data.Hijri, "%d-%d-%d",
			&hijriDate.Year,
			&hijriDate.Month,
			&hijriDate.Day)

		result := hijriDate.ToGregorian().Format("2006-01-02")
		if result != data.Gregorian {
			t.Errorf("%s: want %s got %s\n", data.Hijri, data.Gregorian, result)
		}
	}
}

func Test_Hijri_Bidirectional(t *testing.T) {
	date := time.Date(622, 7, 16, 0, 0, 0, 0, time.UTC)
	for date.Year() <= 2120 {
		// Convert date to hijri
		hijriDate, err := hijri.CreateHijriDate(date, hijri.Default)
		if err != nil {
			date = date.AddDate(0, 0, 1)
			continue
		}

		// Convert back Hijri to Gregorian
		gregorianDate := hijriDate.ToGregorian()

		// Compare original and new gregorian
		strOriginal := date.Format("2006-01-02")
		strGregorian := gregorianDate.Format("2006-01-02")
		strHijri := fmt.Sprintf("%04d-%02d-%02d", hijriDate.Year, hijriDate.Month, hijriDate.Day)

		if strOriginal != strGregorian {
			t.Errorf("Original %s: Hijri %s, Gregorian %s\n",
				strOriginal, strHijri, strGregorian)
		}

		// Increase date
		date = date.AddDate(0, 0, 1)
	}
}
