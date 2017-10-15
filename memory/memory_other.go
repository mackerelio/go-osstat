// +build !linux,!darwin,!windows

package memory

import (
	"fmt"
	"runtime"
)

// Get memory statistics
func Get() (*Memory, error) {
	return nil, fmt.Errorf("memory statistics not implemented for: %s", runtime.GOOS)
}

// Memory represents memory statistics
type Memory struct {
	Total, Used, Cached, Free, Active, Inactive, SwapTotal, SwapUsed, SwapFree uint64
}
