// +build !linux,!darwin

package cpu

import (
	"fmt"
	"runtime"
)

// Get cpu statistics
func Get() (*CPU, error) {
	return nil, fmt.Errorf("cpu statistics not implemented for: %s", runtime.GOOS)
}

// CPU represents cpu statistics
type CPU struct {
	User, System, Idle, Nice, Total uint64
}
