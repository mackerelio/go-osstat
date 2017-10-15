// +build !windows

package main

import (
	"fmt"
	"io"

	"github.com/mackerelio/go-osstat/loadavg"
)

type loadavgGenerator struct {
	loadavg *loadavg.Loadavg
	err     error
}

func (gen *loadavgGenerator) Get() {
	loadavg, err := loadavg.Get()
	if err != nil {
		gen.err = err
		return
	}
	gen.loadavg = loadavg
}

func (gen *loadavgGenerator) Error() error {
	return gen.err
}

func (gen *loadavgGenerator) Print(out io.Writer) {
	loadavg := gen.loadavg
	fmt.Fprintf(out, "loadavg.1\t%f\t-\n", loadavg.Loadavg1)
	fmt.Fprintf(out, "loadavg.5\t%f\t-\n", loadavg.Loadavg5)
	fmt.Fprintf(out, "loadavg.15\t%f\t-\n", loadavg.Loadavg15)
}
