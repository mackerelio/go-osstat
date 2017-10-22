package main

import (
	"github.com/mackerelio/go-osstat/disk"
)

type diskGenerator struct {
	disks []disk.Stats
	err   error
}

func (gen *diskGenerator) Get() {
	gen.disks, gen.err = disk.Get()
}

func (gen *diskGenerator) Error() error {
	return gen.err
}

func (gen *diskGenerator) Print(out chan<- value) {
	for _, disk := range gen.disks {
		out <- value{"disk." + disk.Name + ".reads", disk.ReadsCompleted, "-"}
		out <- value{"disk." + disk.Name + ".writes", disk.WritesCompleted, "-"}
	}
}
