package hijri

import (
	"math"
	"time"
)

// ToGregorian converts Hijri date to Gregorian date.
func ToGregorian(year, month, day int) time.Time {
	// Calculate N days until the end of last year
	Y := year - 1
	yearBy30 := Y / 30
	leftoverYear := Y % 30

	nLeapDays := 0
	for i := 1; i <= leftoverYear; i++ {
		switch i {
		case 2, 5, 7, 10, 13, 16, 18, 21, 24, 26, 29:
			nLeapDays++
		}
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

// ToHijri converts Gregorian date to Hijri.
func ToHijri(date time.Time) (int, int, int) {
	// Calculate Julian Day
	jd := dateToJD(date)

	// Get number of days since 1 Muharram 1H
	flNDays := math.Floor(jd - 1948438.5)
	nDays := int(flNDays)

	// Split days per 30 years, for easier calculation
	yearsBy30 := int(math.Floor(flNDays / 10631.0))
	leftoverDays := nDays % 10631

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
