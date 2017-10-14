// +build !windows

package main

import (
	"fmt"
	"io"

	"github.com/mackerelio/go-osstat/loadavg"
)

type loadavgGenerator struct {
	loadavgs *loadavg.Loadavg
	err      error
}

func (self *loadavgGenerator) Get() {
	loadavgs, err := loadavg.Get()
	if err != nil {
		self.err = err
		return
	}
	self.loadavgs = loadavgs
}

func (self *loadavgGenerator) Error() error {
	return self.err
}

func (self *loadavgGenerator) Print(out io.Writer) {
	loadavgs := self.loadavgs
	fmt.Fprintf(out, "loadavg.1\t%f\t-\n", loadavgs.Loadavg1)
	fmt.Fprintf(out, "loadavg.5\t%f\t-\n", loadavgs.Loadavg5)
	fmt.Fprintf(out, "loadavg.15\t%f\t-\n", loadavgs.Loadavg15)
}
