package hijri_test

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/RadhiFadlillah/go-hijri"
)

var ummAlQuraTests []comparisonTest

func init() {
	// Open test file
	f, err := os.Open("test/ummalqura.csv")
	if err != nil {
		log.Fatalf("failed to open Umm al-Qura test file: %v\n", err)
	}
	defer f.Close()

	// Parse test file
	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to parse Umm al-Qura test file: %v\n", err)
		}

		ummAlQuraTests = append(ummAlQuraTests, comparisonTest{
			Gregorian: record[0],
			Hijri:     record[1],
		})
	}
}

func TestToUmmAlQura(t *testing.T) {
	if len(ummAlQuraTests) == 0 {
		t.Fatal("no tests available for Umm al-Qura")
	}

	for _, item := range ummAlQuraTests {
		date, _ := time.Parse("2006-01-02", item.Gregorian)
		year, month, day, _ := hijri.ToUmmAlQura(date)
		result := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

		if result != item.Hijri {
			t.Errorf("%s: want %s got %s\n", item.Gregorian, item.Hijri, result)
		}
	}
}

func TestFromUmmAlQura(t *testing.T) {
	if len(ummAlQuraTests) == 0 {
		t.Fatal("no tests available for Umm al-Qura")
	}

	for _, item := range ummAlQuraTests {
		var hYear, hMonth, hDay int
		fmt.Sscanf(item.Hijri, "%d-%d-%d", &hYear, &hMonth, &hDay)

		date := hijri.FromUmmAlQura(hYear, hMonth, hDay)
		result := date.Format("2006-01-02")

		if result != item.Gregorian {
			t.Errorf("%s: want %s got %s\n", item.Hijri, item.Gregorian, result)
		}
	}
}

func TestFromToUmmAlQura(t *testing.T) {
	date := time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local)
	for date.Year() < 2031 {
		// Convert date to Umm al-Qura hijri
		hYear, hMonth, hDay, _ := hijri.ToUmmAlQura(date)

		// Convert back Umm al-Qura to Gregorian
		gDate := hijri.FromUmmAlQura(hYear, hMonth, hDay)

		// Compare original and new gregorian
		strHijri := fmt.Sprintf("%04d-%02d-%02d", hYear, hMonth, hDay)
		strOriginal := date.Format("2006-01-02")
		strGregorian := gDate.Format("2006-01-02")

		if strOriginal != strGregorian {
			t.Errorf("%s: hijri %s converted back %s\n",
				strOriginal, strHijri, strGregorian)
		}

		// Increase date
		date = date.AddDate(0, 0, 1)
	}
}
