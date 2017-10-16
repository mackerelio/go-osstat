// +build !linux,!darwin

package network

import (
	"fmt"
	"runtime"
)

// Get network statistics
func Get() ([]NetworkStats, error) {
	return nil, fmt.Errorf("network statistics not implemented for: %s", runtime.GOOS)
}

// NetworkStats represents network statistics
type NetworkStats struct {
	Name             string
	RxBytes, TxBytes uint64
}
