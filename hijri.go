package hijri

import (
	"math"
	"time"
)

// ToHijri converts Gregorian date to standard Hijri date.
func ToHijri(date time.Time) (int, int, int) {
	// We only need the date, so we just set the time to noon
	date = time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, time.UTC)

	// Calculate Julian Day
	jd := dateToJD(date)

	// Get number of days since 1 Muharram 1H
	flNDays := math.Floor(jd - 1948438.5)
	nDays := int(flNDays)

	// Split days per 30 years, for easier calculation
	yearsBy30 := int(math.Floor(flNDays / 10631.0))

	// Get the leftover days
	leftoverDays := nDays - yearsBy30*10631
	if leftoverDays < 0 {
		leftoverDays *= -1
	}

	// From leftover days, calculate year
	var isLeapYear bool
	year := yearsBy30 * 30

	for i := 1; ; i++ {
		year++
		isLeapYear = false

		daysInYear := 354
		switch i {
		case 2, 5, 7, 10, 13, 16, 18, 21, 24, 26, 29:
			isLeapYear = true
			daysInYear = 355
		}

		leftoverDays -= daysInYear
		if leftoverDays <= 0 {
			leftoverDays += daysInYear
			break
		}
	}

	// From leftover days, calculate month and day
	day := 0
	month := 0

	for i := 1; ; i++ {
		month++
		daysInMonth := 29 + i%2
		if isLeapYear && month == 12 {
			daysInMonth = 30
		}

		leftoverDays -= daysInMonth
		if leftoverDays <= 0 {
			day = leftoverDays + daysInMonth
			break
		}
	}

	return year, month, day
}

// FromHijri converts standard Hijri date to Gregorian date.
func FromHijri(year, month, day int) time.Time {
	// Calculate N days until the end of last year
	Y := year - 1
	yearBy30 := Y / 30
	leftoverYear := Y - yearBy30*30
	isNegativeYear := year < 0

	// If the year is negative, for easir calculation, make leftover positive for now
	if isNegativeYear {
		leftoverYear *= -1
	}

	// Count leap days in the leftover years
	nLeapDays := 0
	for i := 1; i <= leftoverYear; i++ {
		switch i {
		case 2, 5, 7, 10, 13, 16, 18, 21, 24, 26, 29:
			nLeapDays++
		}
	}

	// If it was negative, put back the minus
	if isNegativeYear {
		leftoverYear *= -1
		nLeapDays *= -1
	}

	nDaysLastYear := yearBy30*10631 + leftoverYear*354 + nLeapDays

	// Calculate N days in this year, until the end of last month
	M := month - 1
	nDaysLastMonth := 0
	for i := 1; i <= M; i++ {
		if i%2 == 0 {
			nDaysLastMonth += 29
		} else {
			nDaysLastMonth += 30
		}
	}

	// Calculate Julian Days since 1 Muharram 1H
	nDays := nDaysLastYear + nDaysLastMonth + day
	jd := 1948438.5 + float64(nDays)

	return jdToDate(jd)
}
