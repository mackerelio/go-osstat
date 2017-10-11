// +build darwin

package memory

import (
	"bufio"
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
	cmd.Wait()

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
	cmd.Wait()

	memory.SwapTotal = swap.total
	memory.SwapUsed = swap.used
	memory.SwapFree = swap.free
	return memory, nil
}

// Memory represents memory statistics for darwin
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
	freePages        = "Pages free"
	activePages      = "Pages active"
	inactivePages    = "Pages inactive"
	speculativePages = "Pages speculative"
	wiredDownPages   = "Pages wired down"
	purgeablePages   = "Pages purgeable"
	fileBackedPages  = "File-backed pages"
	compressedPages  = "Pages occupied by compressor"
)

// References:
//   - https://support.apple.com/en-us/HT201464#memory
//   - https://developer.apple.com/library/content/documentation/Performance/Conceptual/ManagingMemory/Articles/AboutMemory.html
//   - https://opensource.apple.com/source/system_cmds/system_cmds-790/vm_stat.tproj/
func collectMemoryStats(out io.Reader) (*Memory, error) {
	scanner := bufio.NewScanner(out)
	if !scanner.Scan() { // skip the first line
		return nil, fmt.Errorf("failed to scan output of vm_stat")
	}

	stats := make(map[string]uint64, 22)
	pageSize := uint64(4096)
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.IndexRune(line, ':')
		if i < 0 {
			continue
		}
		val := strings.TrimRight(strings.TrimSpace(line[i+1:]), ".")
		if v, err := strconv.ParseUint(val, 10, 64); err == nil {
			stats[line[:i]] = v * pageSize
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	wired := stats[wiredDownPages]
	compressed := stats[compressedPages]
	cached := stats[purgeablePages] + stats[fileBackedPages]
	active := stats[activePages]
	inactive := stats[inactivePages]
	used := wired + compressed + active + inactive + stats[speculativePages] - cached
	free := stats[freePages]

	return &Memory{
		Total:    used + cached + free,
		Used:     used,
		Cached:   cached,
		Free:     free,
		Active:   active,
		Inactive: inactive,
	}, nil
}

type memorySwap struct {
	total uint64
	free  uint64
	used  uint64
}

func collectSwapStats(out io.Reader) (*memorySwap, error) {
	var total, used, free float64
	_, err := fmt.Fscanf(out, "total = %fM used = %fM free = %fM", &total, &used, &free)
	if err != nil {
		return nil, err
	}
	return &memorySwap{
		total: uint64(total * 1024 * 1024),
		used:  uint64(used * 1024 * 1024),
		free:  uint64(free * 1024 * 1024),
	}, nil
}
