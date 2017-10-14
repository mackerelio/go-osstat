package main

import (
	"fmt"
	"io"

	"github.com/mackerelio/go-osstat/network"
)

type networkGenerator struct {
	networks []network.Network
	err      error
}

func (self *networkGenerator) Get() {
	networks, err := network.Get()
	if err != nil {
		self.err = err
		return
	}
	self.networks = networks
}

func (self *networkGenerator) Error() error {
	return self.err
}

func (self *networkGenerator) Print(out io.Writer) {
	for _, network := range self.networks {
		fmt.Fprintf(out, "network.%s.rx_bytes\t%d\tbytes\n", network.Name, network.RxBytes)
		fmt.Fprintf(out, "network.%s.tx_bytes\t%d\tbytes\n", network.Name, network.TxBytes)
	}
}
