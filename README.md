# Go-Hijri

[![GoDoc](https://godoc.org/github.com/RadhiFadlillah/go-hijri?status.png)](https://godoc.org/github.com/RadhiFadlillah/go-hijri)

Go-Hijri is a Go package for converting Gregorian date to Hijrian date, and vice-versa.

## Usage Examples

```go
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
```

Codes above will give us following results :

```
2019-01-01 AD = 1440-04-23 H
1440-09-01 H = 2019-05-06 AD
```

## Resource

1. Anugraha, R. 2012. _Mekanika Benda Langit_. ([PDF](https://simpan.ugm.ac.id/s/GcxKuyZWn8Rshnn))

## License

Go-Hijri is distributed using [MIT](http://choosealicense.com/licenses/mit/) license.
