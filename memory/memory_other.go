// +build !linux,!darwin,!windows

package memory

import (
	"fmt"
	"runtime"
)

// Get memory statistics
func Get() (*MemoryStats, error) {
	return nil, fmt.Errorf("memory statistics not implemented for: %s", runtime.GOOS)
}

// MemoryStats represents memory statistics
type MemoryStats struct {
	Total, Used, Cached, Free, Active, Inactive, SwapTotal, SwapUsed, SwapFree uint64
}
