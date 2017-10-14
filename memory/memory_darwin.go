// +build darwin

package memory

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

// Get memory statistics
func Get() (*Memory, error) {
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

	cmd = exec.Command("sysctl", "-n", "vm.swapusage")
	out, err = cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	swap, err := collectSwapStats(out)
	if err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	memory.SwapTotal = swap.total
	memory.SwapUsed = swap.used
	memory.SwapFree = swap.free
	return memory, nil
}

// Memory represents memory statistics for darwin
type Memory struct {
	Total, Used, Cached, Free, Active, Inactive, SwapTotal, SwapUsed, SwapFree uint64
}

// References:
//   - https://support.apple.com/en-us/HT201464#memory
//   - https://developer.apple.com/library/content/documentation/Performance/Conceptual/ManagingMemory/Articles/AboutMemory.html
//   - https://opensource.apple.com/source/system_cmds/system_cmds-790/vm_stat.tproj/
func collectMemoryStats(out io.Reader) (*Memory, error) {
	scanner := bufio.NewScanner(out)
	if !scanner.Scan() { // skip the first line
		return nil, fmt.Errorf("failed to scan output of vm_stat")
	}

	var memory Memory
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
		return nil, err
	}

	memory.Cached = purgeable + fileBacked
	memory.Used = wired + compressed + memory.Active + memory.Inactive + speculative - memory.Cached
	memory.Total = memory.Used + memory.Cached + memory.Free
	return &memory, nil
}

type memorySwap struct {
	total, free, used uint64
}

func collectSwapStats(out io.Reader) (*memorySwap, error) {
	var total, used, free float64
	cnt, err := fmt.Fscanf(out, "total = %fM used = %fM free = %fM", &total, &used, &free)
	if err != nil {
		return nil, fmt.Errorf("failed to scan output of sysctl -n vmswapusage: %s", err)
	}
	if cnt != 3 {
		return nil, errors.New("failed to parse output of 'sysctl -n vm.swapusage'")
	}
	return &memorySwap{
		total: uint64(total * 1024 * 1024),
		used:  uint64(used * 1024 * 1024),
		free:  uint64(free * 1024 * 1024),
	}, nil
}
