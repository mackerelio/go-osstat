package main

import (
	"github.com/mackerelio/go-osstat/network"
)

type networkGenerator struct {
	networks []network.Stats
	err      error
}

func (gen *networkGenerator) Get() {
	gen.networks, gen.err = network.Get()
}

func (gen *networkGenerator) Error() error {
	return gen.err
}

func (gen *networkGenerator) Print(out chan<- value) {
	for _, network := range gen.networks {
		out <- value{"network." + network.Name + ".rx_bytes", network.RxBytes, "bytes"}
		out <- value{"network." + network.Name + ".tx_bytes", network.TxBytes, "bytes"}
	}
}
