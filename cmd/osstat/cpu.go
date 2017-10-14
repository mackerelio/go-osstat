package main

import (
	"fmt"
	"io"

	"github.com/mackerelio/go-osstat/cpu"
)

type cpuGenerator struct {
	cpu *cpu.Cpu
	err error
}

func (self *cpuGenerator) Get() {
	cpu, err := cpu.Get()
	if err != nil {
		self.err = err
		return
	}
	self.cpu = cpu
}

func (self *cpuGenerator) Error() error {
	return self.err
}

func (self *cpuGenerator) Print(out io.Writer) {
	cpu := self.cpu
	fmt.Fprintf(out, "cpu.user\t%f\t%%\n", cpu.User)
	fmt.Fprintf(out, "cpu.system\t%f\t%%\n", cpu.System)
	fmt.Fprintf(out, "cpu.idle\t%f\t%%\n", cpu.Idle)
}
