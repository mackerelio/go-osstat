// +build linux

package memory

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// Get memory statistics
func Get() (*Memory, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	memory, err := collectMemoryStats(file)
	if err != nil {
		return nil, err
	}
	return memory, nil
}

// Memory represents memory statistics for linux
type Memory struct {
	Total     uint64
	Used      uint64
	Cached    uint64
	Free      uint64
	Active    uint64
	Inactive  uint64
	SwapTotal uint64
	SwapUsed  uint64
	SwapFree  uint64
}

const (
	memTotal      = "MemTotal"
	memFree       = "MemFree"
	memAvailable  = "MemAvailable"
	memBuffers    = "Buffers"
	memCached     = "Cached"
	memActive     = "Active"
	memInactive   = "Inactive"
	memSwapCached = "SwapCached"
	memSwapTotal  = "SwapTotal"
	memSwapFree   = "SwapFree"
)

func collectMemoryStats(out io.Reader) (*Memory, error) {
	scanner := bufio.NewScanner(out)
	stats := make(map[string]uint64, 10)
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.IndexRune(line, ':')
		if i < 0 {
			continue
		}
		switch line[:i] {
		case memTotal, memFree, memAvailable, memBuffers, memCached,
			memActive, memInactive, memSwapCached, memSwapTotal, memSwapFree:
			val := strings.TrimSpace(strings.TrimRight(line[i+1:], "kB"))
			if v, err := strconv.ParseUint(val, 10, 64); err == nil {
				stats[line[:i]] = v * 1024
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Memory{
		Total:     stats[memTotal],
		Used:      stats[memTotal] - stats[memFree] - stats[memBuffers] - stats[memCached],
		Cached:    stats[memCached],
		Free:      stats[memFree],
		Active:    stats[memActive],
		Inactive:  stats[memInactive],
		SwapTotal: stats[memSwapTotal],
		SwapUsed:  stats[memSwapTotal] - stats[memSwapFree],
		SwapFree:  stats[memSwapFree],
	}, nil
}
