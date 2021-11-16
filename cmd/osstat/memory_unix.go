//go:build !windows
// +build !windows

package main

func (gen *memoryGenerator) Print(out chan<- value) {
	memory := gen.memory
	out <- value{"memory.total", memory.Total, "bytes"}
	out <- value{"memory.used", memory.Used, "bytes"}
	out <- value{"memory.cached", memory.Cached, "bytes"}
	out <- value{"memory.free", memory.Free, "bytes"}
	out <- value{"memory.active", memory.Active, "bytes"}
	out <- value{"memory.inactive", memory.Inactive, "bytes"}
	out <- value{"memory.swap_total", memory.SwapTotal, "bytes"}
	out <- value{"memory.swap_used", memory.SwapUsed, "bytes"}
	out <- value{"memory.swap_free", memory.SwapFree, "bytes"}
}
