// +build freebsd

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
func Get() (*Stats, error) {
	cmd := exec.Command("top", "-b", "-n", "1")
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
	return memory, nil
}

// Stats represents memory statistics for freebsd
type Stats struct {
	Total, Used, Buffers, Free, Active, Inactive, Wired,
	SwapTotal, SwapFree uint64
}

func collectMemoryStats(out io.Reader) (*Stats, error) {
	scanner := bufio.NewScanner(out)

	var memory Stats
	memStats := map[string]*uint64{
		"MemBuf":    &memory.Buffers,
		"MemFree":   &memory.Free,
		"MemActive": &memory.Active,
		"MemInact":  &memory.Inactive,
		"MemWired":  &memory.Wired,
		"SwapTotal": &memory.SwapTotal,
		"SwapFree":  &memory.SwapFree,
	}

	var cnt int
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Mem:") && !strings.HasPrefix(line, "Swap:") {
			continue
		}
		cnt += 1
		i := strings.IndexRune(line, ':')
		prefix := line[:i]
		stats := strings.Split(line[i+1:], ",")
		for _, stat := range stats {
			cs := strings.Fields(stat)
			if len(cs) != 2 {
				continue
			}
			if ptr := memStats[prefix+cs[1]]; ptr != nil {
				if val, err := parseValue(cs[0]); err == nil {
					*ptr = val
				}
			}
		}
		if cnt == 2 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for top -b -n 1: %s", err)
	}

	return &memory, nil
}

func parseValue(valStr string) (uint64, error) {
	if len(valStr) < 1 {
		return 0, errors.New("empty value")
	}
	var unit uint64
	switch valStr[len(valStr)-1] {
	case 'T':
		unit = 1024 * 1024 * 1024 * 1024
	case 'G':
		unit = 1024 * 1024 * 1024
	case 'M':
		unit = 1024 * 1024
	case 'K':
		unit = 1024
	default:
		unit = 1
	}
	val, err := strconv.ParseUint(valStr[:len(valStr)-1], 10, 64)
	return val * unit, err
}
