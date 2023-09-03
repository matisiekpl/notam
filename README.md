# NOTAM
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
