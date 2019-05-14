package hijri

import (
	"time"

	dec "github.com/shopspring/decimal"
)

func dateToJD(date time.Time) float64 {
	// Convert to UTC
	date = date.UTC()

	// Prepare variables for calculating
	Y := int64(date.Year())
	M := int64(date.Month())
	D := int64(date.Day())
	H := int64(date.Hour())
	m := int64(date.Minute())
	s := int64(date.Second())

	// If year is before 4713 B.C, stop
	if Y < -4712 {
		return 0
	}

	// If date is in blank days, stop
	endOfJulian := time.Date(1582, 10, 4, 23, 59, 59, 0, time.UTC)
	startOfGregorian := time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC)
	if date.After(endOfJulian) && date.Before(startOfGregorian) {
		return 0
	}

	// If month <= 2, change year and month
	if M <= 2 {
		M += 12
		Y--
	}

	// Check whether date is gregorian or julian
	constant := dec.Zero
	if date.After(endOfJulian) {
		temp := dec.New(Y, -2).Floor()
		constant = dec.New(2, 0).
			Add(temp.Div(dec.New(4, 0)).Floor()).
			Sub(temp)
	}

	// Calculate julian day
	yearToDays := dec.New(Y, 0).
		Mul(dec.NewFromFloat(365.25)).
		Floor()

	monthToDays := dec.New(M+1, 0).
		Mul(dec.NewFromFloat(30.6001)).
		Floor()

	timeToSeconds := H*3600 + m*60 + s
	timeToDays := dec.New(timeToSeconds, 0).
		Div(dec.New(86400, 0))

	julianDay, _ := dec.NewFromFloat(1720994.5).
		Add(yearToDays).
		Add(monthToDays).
		Add(constant).
		Add(dec.New(D, 0)).
		Add(timeToDays).
		Float64()

	return julianDay
}

func jdToDate(jd float64) time.Time {
	// Prepare variables for calculating
	jd1 := dec.NewFromFloat(jd).Add(dec.NewFromFloat(0.5))
	z := jd1.Floor()
	f := jd1.Sub(z)

	a := z
	if z.GreaterThanOrEqual(dec.New(2299161, 0)) {
		aa := z.Sub(dec.NewFromFloat(1867216.25)).
			Div(dec.NewFromFloat(36524.25)).
			Floor()
		aaBy4 := aa.Div(dec.New(4, 0)).Floor()
		a = z.Add(dec.New(1, 0)).Add(aa).Sub(aaBy4)
	}

	b := a.Add(dec.New(1524, 0))
	c := b.Sub(dec.NewFromFloat(122.1)).
		Div(dec.NewFromFloat(365.25)).
		Floor()
	d := c.Mul(dec.NewFromFloat(365.25)).Floor()
	e := b.Sub(d).Div(dec.NewFromFloat(30.6001)).Floor()

	// Calculate day with its time
	dayTime := b.Sub(d).
		Sub(e.Mul(dec.NewFromFloat(30.6001)).Floor()).
		Add(f)
	day := dayTime.Floor()

	// Calculate time
	seconds := dayTime.Sub(day).Mul(dec.New(24*60*60, 0))
	hour := seconds.Div(dec.New(3600, 0)).Floor()
	min := seconds.Sub(hour.Mul(dec.New(3600, 0))).
		Div(dec.New(60, 0)).
		Floor()
	sec := seconds.Sub(hour.Mul(dec.New(3600, 0))).
		Sub(min.Mul(dec.New(60, 0))).
		Floor()

	// Calculate month
	var month dec.Decimal
	if e.LessThan(dec.New(14, 0)) {
		month = e.Sub(dec.New(1, 0))
	} else {
		month = e.Sub(dec.New(13, 0))
	}

	// Calculate year
	var year dec.Decimal
	if month.GreaterThan(dec.New(2, 0)) {
		year = c.Sub(dec.New(4716, 0))
	} else {
		year = c.Sub(dec.New(4715, 0))
	}

	// Create date
	intYear := int(year.IntPart())
	intMonth := int(month.IntPart())
	intDay := int(day.IntPart())
	intHour := int(hour.IntPart())
	intMin := int(min.IntPart())
	intSec := int(sec.IntPart())

	return time.Date(intYear, time.Month(intMonth), intDay,
		intHour, intMin, intSec, 0, time.UTC)
}
