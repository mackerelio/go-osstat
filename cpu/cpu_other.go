// +build !linux,!darwin

package cpu

import (
	"fmt"
	"runtime"
)

// Get cpu statistics
func Get() (*Cpu, error) {
	return nil, fmt.Errorf("cpu statistics not implemented for: %s", runtime.GOOS)
}

// Cpu represents cpu statistics
type Cpu struct {
	User, System, Idle, Nice, Total uint64
}
