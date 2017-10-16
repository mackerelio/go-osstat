// +build !linux,!darwin

package cpu

import (
	"fmt"
	"runtime"
)

// Get cpu statistics
func Get() (*CPUStats, error) {
	return nil, fmt.Errorf("cpu statistics not implemented for: %s", runtime.GOOS)
}

// CPUStats represents cpu statistics
type CPUStats struct {
	User, System, Idle, Nice, Total uint64
}
