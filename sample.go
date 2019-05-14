// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/RadhiFadlillah/go-hijri"
)

func main() {
	// 1 January 2019 to Hijri
	src := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
	y, m, d := hijri.ToHijri(src)
	fmt.Printf("%s AD = %04d-%02d-%02d H\n", src.Format("2006-01-02"), y, m, d)

	// 1 Ramadhan 1440 to Gregorian
	ramadhan := hijri.ToGregorian(1440, 9, 1)
	fmt.Printf("1440-09-01 H = %s AD\n", ramadhan.Format("2006-01-02"))
}
