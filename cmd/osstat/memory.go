package main

import (
	"fmt"
	"io"

	"github.com/mackerelio/go-osstat/memory"
)

type memoryGenerator struct {
	memory *memory.Memory
	err    error
}

func (gen *memoryGenerator) Get() {
	memory, err := memory.Get()
	if err != nil {
		gen.err = err
		return
	}
	gen.memory = memory
}

func (gen *memoryGenerator) Error() error {
	return gen.err
}

func (gen *memoryGenerator) Print(out io.Writer) {
	memory := gen.memory
	fmt.Fprintf(out, "memory.total\t%d\tbytes\n", memory.Total)
	fmt.Fprintf(out, "memory.used\t%d\tbytes\n", memory.Used)
	fmt.Fprintf(out, "memory.cached\t%d\tbytes\n", memory.Cached)
	fmt.Fprintf(out, "memory.free\t%d\tbytes\n", memory.Free)
	fmt.Fprintf(out, "memory.active\t%d\tbytes\n", memory.Active)
	fmt.Fprintf(out, "memory.inactive\t%d\tbytes\n", memory.Inactive)
	fmt.Fprintf(out, "memory.swap_total\t%d\tbytes\n", memory.SwapTotal)
	fmt.Fprintf(out, "memory.swap_used\t%d\tbytes\n", memory.SwapUsed)
	fmt.Fprintf(out, "memory.swap_free\t%d\tbytes\n", memory.SwapFree)
}
