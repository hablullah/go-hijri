package hijri_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RadhiFadlillah/go-hijri"
)

var ummAlQuraTestData []TestData

func init() {
	var err error
	ummAlQuraTestData, err = generateTestData("test/ummalqura.csv")
	if err != nil {
		panic(err)
	}
}

func Test_UmmAlQura_ConvertDate(t *testing.T) {
	if len(ummAlQuraTestData) == 0 {
		t.Fatal("no tests available for Umm al-Qura")
	}

	for _, data := range ummAlQuraTestData {
		gregorianDate, _ := time.Parse("2006-01-02", data.Gregorian)
		ummAlQuraDate, _ := hijri.CreateUmmAlQuraDate(gregorianDate)
		strUmmAlQuraDate := fmt.Sprintf("%04d-%02d-%02d",
			ummAlQuraDate.Year,
			ummAlQuraDate.Month,
			ummAlQuraDate.Day)

		if strUmmAlQuraDate != data.Hijri {
			t.Errorf("%s: want %s got %s\n", data.Gregorian, data.Hijri, strUmmAlQuraDate)
		}
	}
}

func Test_UmmAlQura_ToGregorian(t *testing.T) {
	if len(ummAlQuraTestData) == 0 {
		t.Fatal("no tests available for Umm al-Qura")
	}

	for _, data := range ummAlQuraTestData {
		var ummAlQuraDate hijri.UmmAlQuraDate
		fmt.Sscanf(data.Hijri, "%d-%d-%d",
			&ummAlQuraDate.Year,
			&ummAlQuraDate.Month,
			&ummAlQuraDate.Day)

		result := ummAlQuraDate.ToGregorian().Format("2006-01-02")
		if result != data.Gregorian {
			t.Errorf("%s: want %s got %s\n", data.Hijri, data.Gregorian, result)
		}
	}
}

func Test_UmmAlQura_Bidirectional(t *testing.T) {
	date := time.Date(1937, 3, 14, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(2077, 11, 17, 0, 0, 0, 0, time.UTC)
	for date.Before(maxDate) {
		// Convert date to Umm al-Qura
		ummAlQuraDate, err := hijri.CreateUmmAlQuraDate(date)
		if err != nil {
			date = date.AddDate(0, 0, 1)
			continue
		}

		// Convert back Umm al-Qura to Gregorian
		gregorianDate := ummAlQuraDate.ToGregorian()

		// Compare original and new gregorian
		strOriginal := date.Format("2006-01-02")
		strGregorian := gregorianDate.Format("2006-01-02")
		strHijri := fmt.Sprintf("%04d-%02d-%02d", ummAlQuraDate.Year, ummAlQuraDate.Month, ummAlQuraDate.Day)

		if strOriginal != strGregorian {
			t.Errorf("Original %s: Hijri %s, Gregorian %s\n",
				strOriginal, strHijri, strGregorian)
		}

		// Increase date
		date = date.AddDate(0, 0, 1)
	}
}
