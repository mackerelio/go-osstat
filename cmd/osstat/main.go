// osstat shows os system metric statistics.
//
// Usage:
//
//	osstat
//
package main

import (
	"fmt"
	"os"
)

var name = "osstat"

func main() {
	if errs := run(os.Args[1:], os.Stdout); errs != nil {
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
		}
		os.Exit(1)
	}
}
