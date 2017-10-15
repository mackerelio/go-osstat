// +build !linux,!darwin

package network

import (
	"fmt"
	"runtime"
)

// Get network statistics
func Get() ([]Network, error) {
	return nil, fmt.Errorf("network statistics not implemented for: %s", runtime.GOOS)
}

// Network represents network statistics
type Network struct {
	Name             string
	RxBytes, TxBytes uint64
}
