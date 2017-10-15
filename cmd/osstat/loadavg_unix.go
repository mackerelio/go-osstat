// +build !windows

package main

import (
	"fmt"
	"io"

	"github.com/mackerelio/go-osstat/loadavg"
)

type loadavgGenerator struct {
	loadavg *loadavg.Loadavg
	err     error
}

func (self *loadavgGenerator) Get() {
	loadavg, err := loadavg.Get()
	if err != nil {
		self.err = err
		return
	}
	self.loadavg = loadavg
}

func (self *loadavgGenerator) Error() error {
	return self.err
}

func (self *loadavgGenerator) Print(out io.Writer) {
	loadavg := self.loadavg
	fmt.Fprintf(out, "loadavg.1\t%f\t-\n", loadavg.Loadavg1)
	fmt.Fprintf(out, "loadavg.5\t%f\t-\n", loadavg.Loadavg5)
	fmt.Fprintf(out, "loadavg.15\t%f\t-\n", loadavg.Loadavg15)
}
