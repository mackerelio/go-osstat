package main

import (
	"fmt"
	"io"

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

func (gen *diskGenerator) Print(out io.Writer) {
	for _, disk := range gen.disks {
		fmt.Fprintf(out, "disk.%s.reads\t%d\t-\n", disk.Name, disk.ReadsCompleted)
		fmt.Fprintf(out, "disk.%s.writes\t%d\t-\n", disk.Name, disk.WritesCompleted)
	}
}
