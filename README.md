# addy
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/deelawn/addy)
[![GoReportCard](https://goreportcard.com/badge/github.com/nanomsg/mangos)](https://goreportcard.com/report/github.com/deelawn/addy)
Basic US Address Normalization

---

### Usage
```go
package main

import (
	"fmt"

	"github.com/deelawn/addy"
)

func main() {

	// address = "405 NW BEN FRANKLIN PKWY"
	// count = 2
	address, count := addy.NormalizeAddress1("405 northwest Ben Franklin parkway")
	fmt.Println(address, count)

	// address = "25TH FL, BOX 9"
	// count = 2
	addres, count = addy.NormalizeAddress2("25th floor, box 9")
	fmt.Println(address, count)
}
```
