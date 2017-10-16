package main

import (
	"fmt"
	"io"

	"github.com/mackerelio/go-osstat/network"
)

type networkGenerator struct {
	networks []network.NetworkStats
	err      error
}

func (gen *networkGenerator) Get() {
	networks, err := network.Get()
	if err != nil {
		gen.err = err
		return
	}
	gen.networks = networks
}

func (gen *networkGenerator) Error() error {
	return gen.err
}

func (gen *networkGenerator) Print(out io.Writer) {
	for _, network := range gen.networks {
		fmt.Fprintf(out, "network.%s.rx_bytes\t%d\tbytes\n", network.Name, network.RxBytes)
		fmt.Fprintf(out, "network.%s.tx_bytes\t%d\tbytes\n", network.Name, network.TxBytes)
	}
}
