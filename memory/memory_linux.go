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
	Total, Used, Cached, Free, Active, Inactive, SwapTotal, SwapUsed, SwapCached, SwapFree uint64
}

func collectMemoryStats(out io.Reader) (*Memory, error) {
	scanner := bufio.NewScanner(out)
	var memory Memory
	var buffers uint64
	memStats := map[string]*uint64{
		"MemTotal":   &memory.Total,
		"MemFree":    &memory.Free,
		"Buffers":    &buffers,
		"Cached":     &memory.Cached,
		"Active":     &memory.Active,
		"Inactive":   &memory.Inactive,
		"SwapCached": &memory.SwapCached,
		"SwapTotal":  &memory.SwapTotal,
		"SwapFree":   &memory.SwapFree,
	}
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.IndexRune(line, ':')
		if i < 0 {
			continue
		}
		if ptr := memStats[line[:i]]; ptr != nil {
			val := strings.TrimSpace(strings.TrimRight(line[i+1:], "kB"))
			if v, err := strconv.ParseUint(val, 10, 64); err == nil {
				*ptr = v * 1024
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	memory.SwapUsed = memory.SwapTotal - memory.SwapFree
	memory.Used = memory.Total - memory.Free - buffers - memory.Cached
	return &memory, nil
}
