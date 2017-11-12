package main

import (
	"fmt"
	"time"

	"github.com/mackerelio/go-osstat/uptime"
)

type uptimeGenerator struct {
	uptime time.Duration
	err    error
}

func (gen *uptimeGenerator) Get() {
	gen.uptime, gen.err = uptime.Get()
}

func (gen *uptimeGenerator) Error() error {
	return gen.err
}

func (gen *uptimeGenerator) Print(out chan<- value) {
	uptime := gen.uptime
	out <- value{"uptime", fmt.Sprintf("%f", float64(uptime.Nanoseconds())/1e9), "seconds"}
}
