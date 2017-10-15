package main

import (
	"fmt"
	"io"
)

func (gen *memoryGenerator) Print(out io.Writer) {
	memory := gen.memory
	fmt.Fprintf(out, "memory.total\t%d\tbytes\n", memory.Total)
	fmt.Fprintf(out, "memory.used\t%d\tbytes\n", memory.Used)
	fmt.Fprintf(out, "memory.free\t%d\tbytes\n", memory.Free)
	fmt.Fprintf(out, "memory.page_file_total\t%d\tbytes\n", memory.PageFileTotal)
	fmt.Fprintf(out, "memory.page_file_free\t%d\tbytes\n", memory.PageFileFree)
	fmt.Fprintf(out, "memory.virtual_total\t%d\tbytes\n", memory.VirtualTotal)
	fmt.Fprintf(out, "memory.virtual_free\t%d\tbytes\n", memory.VirtualFree)
}
