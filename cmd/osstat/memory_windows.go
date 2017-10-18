package main

func (gen *memoryGenerator) Print(out chan<- value) {
	memory := gen.memory
	out <- value{"memory.total", memory.Total, "bytes"}
	out <- value{"memory.used", memory.Used, "bytes"}
	out <- value{"memory.free", memory.Free, "bytes"}
	out <- value{"memory.page_file_total", memory.PageFileTotal, "bytes"}
	out <- value{"memory.page_file_free", memory.PageFileFree, "bytes"}
	out <- value{"memory.virtual_total", memory.VirtualTotal, "bytes"}
	out <- value{"memory.virtual_free", memory.VirtualFree, "bytes"}
}
