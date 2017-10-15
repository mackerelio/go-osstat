package main

import (
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
