package main

import (
	"github.com/mackerelio/go-osstat/disk"
)

type diskGenerator struct {
	disks []disk.DiskStats
	err   error
}

func (gen *diskGenerator) Get() {
	disks, err := disk.Get()
	if err != nil {
		gen.err = err
		return
	}
	gen.disks = disks
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
