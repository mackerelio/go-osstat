package main

import (
	"fmt"
	"os"
)

var name = "osstat"
var version = "v0.0.0"
var description = "show os system metric statistics"
var author = ""

func main() {
	if errs := run(os.Args[1:], os.Stdout); errs != nil {
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
		}
		os.Exit(1)
	}
}
