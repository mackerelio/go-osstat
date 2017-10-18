// +build !windows

package main

import (
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

func (gen *loadavgGenerator) Print(out chan<- value) {
	loadavg := gen.loadavg
	out <- value{"loadavg.1", loadavg.Loadavg1, "-"}
	out <- value{"loadavg.5", loadavg.Loadavg5, "-"}
	out <- value{"loadavg.15", loadavg.Loadavg15, "-"}
}
