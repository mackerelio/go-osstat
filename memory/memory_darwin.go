// +build darwin

package memory

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// Get memory statistics
func Get() (*MemoryStats, error) {
	cmd := exec.Command("vm_stat")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	memory, err := collectMemoryStats(out)
	if err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	ret, err := syscall.Sysctl("vm.swapusage")
	if err != nil {
		return nil, fmt.Errorf("failed in sysctl vm.swapusage: %s", err)
	}
	swap, err := collectSwapStats(strings.NewReader(ret))
	if err != nil {
		return nil, err
	}
	memory.SwapTotal = swap.Total
	memory.SwapUsed = swap.Used
	memory.SwapFree = swap.Avail

	return memory, nil
}

// MemoryStats represents memory statistics for darwin
type MemoryStats struct {
	Total, Used, Cached, Free, Active, Inactive, SwapTotal, SwapUsed, SwapFree uint64
}

// References:
//   - https://support.apple.com/en-us/HT201464#memory
//   - https://developer.apple.com/library/content/documentation/Performance/Conceptual/ManagingMemoryStats/Articles/AboutMemoryStats.html
//   - https://opensource.apple.com/source/system_cmds/system_cmds-790/vm_stat.tproj/
func collectMemoryStats(out io.Reader) (*MemoryStats, error) {
	scanner := bufio.NewScanner(out)
	if !scanner.Scan() { // skip the first line
		return nil, fmt.Errorf("failed to scan output of vm_stat")
	}

	var memory MemoryStats
	var speculative, wired, purgeable, fileBacked, compressed uint64
	memStats := map[string]*uint64{
		"Pages free":                   &memory.Free,
		"Pages active":                 &memory.Active,
		"Pages inactive":               &memory.Inactive,
		"Pages speculative":            &speculative,
		"Pages wired down":             &wired,
		"Pages purgeable":              &purgeable,
		"File-backed pages":            &fileBacked,
		"Pages occupied by compressor": &compressed,
	}
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.IndexRune(line, ':')
		if i < 0 {
			continue
		}
		if ptr := memStats[line[:i]]; ptr != nil {
			val := strings.TrimRight(strings.TrimSpace(line[i+1:]), ".")
			if v, err := strconv.ParseUint(val, 10, 64); err == nil {
				*ptr = v * 4096
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for vm_stat: %s", err)
	}

	memory.Cached = purgeable + fileBacked
	memory.Used = wired + compressed + memory.Active + memory.Inactive + speculative - memory.Cached
	memory.Total = memory.Used + memory.Cached + memory.Free
	return &memory, nil
}

// xsw_usage in sysctl.h
type swapUsage struct {
	Total     uint64
	Avail     uint64
	Used      uint64
	Pagesize  int32
	Encrypted bool
}

func collectSwapStats(out io.Reader) (*swapUsage, error) {
	var swap swapUsage
	err := binary.Read(out, binary.LittleEndian, &swap)
	if err != nil {
		return nil, fmt.Errorf("failed to read result of sysctl vm.swapusage: %s", err)
	}
	return &swap, nil
}
