//go:build !windows
// +build !windows

package main

import (
	"github.com/mackerelio/go-osstat/loadavg"
)

type loadavgGenerator struct {
	loadavg *loadavg.Stats
	err     error
}

func (gen *loadavgGenerator) Get() {
	gen.loadavg, gen.err = loadavg.Get()
}

func (gen *loadavgGenerator) Error() error {
	return gen.err
}

func (gen *loadavgGenerator) Print(out chan<- value) {
	loadavg := gen.loadavg
	out <- value{"loadavg.1m", loadavg.Loadavg1, "-"}
	out <- value{"loadavg.5m", loadavg.Loadavg5, "-"}
	out <- value{"loadavg.15m", loadavg.Loadavg15, "-"}
}
