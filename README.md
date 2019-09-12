# Go-Hijri

[![GoDoc](https://godoc.org/github.com/RadhiFadlillah/go-hijri?status.png)](https://godoc.org/github.com/RadhiFadlillah/go-hijri)

Go-Hijri is a Go package for converting Gregorian date to Hijrian date and vice-versa. There are two supported Hijrian calendar :

- the standard Hijri date which calculated based on standard astronomical method,
- the Umm al-Qura calendar which used by Saudi Arabia and most Islamic country.

## Usage Examples

```go
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

```

Codes above will give us following results :

```
2019-01-01 AD = 1440-04-23 H (standard)
2019-01-01 AD = 1440-04-25 H (umm al-qura)
1440-09-01 H = 2019-05-06 AD
```

## Resource

1. Anugraha, R. 2012. _Mekanika Benda Langit_. ([PDF](https://simpan.ugm.ac.id/s/GcxKuyZWn8Rshnn))
2. Van Gent, R. H. 2019. _The Umm al-Qura Calendar of Saudi Arabia_. ([Website](https://www.staff.science.uu.nl/~gent0113/islam/ummalqura.htm))
3. Strous, Dr. 2019. _Astronomy Answers: Julian Days Number_. ([Website](https://www.aa.quae.nl/en/reken/juliaansedag.html))

## License

Go-Hijri is distributed using [MIT](http://choosealicense.com/licenses/mit/) license.
