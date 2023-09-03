# NOTAM
<a href="https://pkg.go.dev/github.com/matisiekpl/notam"><img src="https://pkg.go.dev/badge/github.com/matisiekpl/notam.svg" alt="Go Reference"></a>

Fetch NOTAM for given ICAO code from www.notams.faa.gov

## Installation 
```bash
go get github.com/matisiekpl/notam
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/matisiekpl/notam"
	"log"
)

func main() {
	// Rzeszow-Jasionka Airport, Poland
	items, err := notam.Fetch("EPRZ")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}

```
