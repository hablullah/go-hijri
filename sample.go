// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/RadhiFadlillah/go-hijri"
)

func main() {
	// 1 January 2019 to standard Hijri
	newYear := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
	y, m, d := hijri.ToHijri(newYear)
	fmt.Printf("%s AD = %04d-%02d-%02d H (standard)\n", newYear.Format("2006-01-02"), y, m, d)

	// 1 January 2019 to Umm al-Qura calendar
	y, m, d, _ = hijri.ToUmmAlQura(newYear)
	fmt.Printf("%s AD = %04d-%02d-%02d H (umm al-qura)\n", newYear.Format("2006-01-02"), y, m, d)

	// 1 Ramadhan 1440 to Gregorian
	ramadhan := hijri.ToGregorian(1440, 9, 1)
	fmt.Printf("1440-09-01 H = %s AD\n", ramadhan.Format("2006-01-02"))
}
