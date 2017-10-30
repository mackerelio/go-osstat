// +build freebsd

package memory

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/sys/unix"
)

// Get memory statistics
func Get() (*Stats, error) {
	memory, err := collectMemoryStats()
	if err != nil {
		return nil, err
	}

	// Maybe this is incorrect... needs more research on memory statistics...
	memory.Used = memory.Total - memory.Free - memory.Cached - memory.Inactive
	memory.SwapUsed = memory.SwapTotal - memory.SwapFree

	return memory, nil
}

// Stats represents memory statistics for freebsd
type Stats struct {
	Total, Used, Cached, Free, Active, Inactive, Wired,
	SwapTotal, SwapUsed, SwapFree uint64
}

type memStat struct {
	name  string
	ptr   *uint64
	scale *uint64
}

func collectMemoryStats() (*Stats, error) {
	var pageSize uint64
	one := uint64(1)

	var memory Stats
	memStats := []memStat{
		{"vm.stats.vm.v_page_size", &pageSize, &one},
		{"hw.physmem", &memory.Total, &one},
		{"vm.stats.vm.v_cache_count", &memory.Cached, &pageSize},
		{"vm.stats.vm.v_free_count", &memory.Free, &pageSize},
		{"vm.stats.vm.v_active_count", &memory.Active, &pageSize},
		{"vm.stats.vm.v_inactive_count", &memory.Inactive, &pageSize},
		{"vm.stats.vm.v_wire_count", &memory.Wired, &pageSize},
		{"vm.swap_total", &memory.SwapTotal, &one},
	}
	// TODO: swap_free

	for _, stat := range memStats {
		ret, err := unix.SysctlRaw(stat.name)
		if err != nil {
			return nil, fmt.Errorf("failed in sysctl %s: %s", stat.name, err)
		}
		if len(ret) == 8 {
			*stat.ptr = binary.LittleEndian.Uint64(ret) * *stat.scale
		} else if len(ret) == 4 {
			*stat.ptr = uint64(binary.LittleEndian.Uint32(ret)) * *stat.scale
		} else {
			return nil, fmt.Errorf("failed in sysctl %s: %s", stat.name, err)
		}
	}

	return &memory, nil
}
