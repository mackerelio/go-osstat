// +build !darwin darwin,cgo

package main

import (
	"fmt"
	"io"

	"github.com/mackerelio/go-osstat/cpu"
)

type cpuGenerator struct {
	cpu *cpu.CPUStats
	err error
}

func (gen *cpuGenerator) Get() {
	cpu, err := cpu.Get()
	if err != nil {
		gen.err = err
		return
	}
	gen.cpu = cpu
}

func (gen *cpuGenerator) Error() error {
	return gen.err
}

func (gen *cpuGenerator) Print(out io.Writer) {
	cpu := gen.cpu
	fmt.Fprintf(out, "cpu.user\t%d\t-\n", cpu.User)
	fmt.Fprintf(out, "cpu.system\t%d\t-\n", cpu.System)
	fmt.Fprintf(out, "cpu.idle\t%d\t-\n", cpu.Idle)
	fmt.Fprintf(out, "cpu.total\t%d\t-\n", cpu.Total)
}
