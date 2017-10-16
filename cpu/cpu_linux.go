// +build linux

package cpu

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Get cpu statistics
func Get() (*CPU, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cpu, err := collectCPUStats(file)
	if err != nil {
		return nil, err
	}
	return cpu, nil
}

// CPU represents cpu statistics for linux
type CPU struct {
	User, Nice, System, Idle, Iowait, Irq, Softirq, Steal, Guest, GuestNice uint64
	Total, CPUCount                                                         uint64
}

type cpuStat struct {
	name string
	ptr  *uint64
}

func collectCPUStats(out io.Reader) (*CPU, error) {
	scanner := bufio.NewScanner(out)
	var cpu CPU

	cpuStats := []cpuStat{
		{"user", &cpu.User},
		{"nice", &cpu.Nice},
		{"system", &cpu.System},
		{"idle", &cpu.Idle},
		{"iowait", &cpu.Iowait},
		{"irq", &cpu.Irq},
		{"softirq", &cpu.Softirq},
		{"steal", &cpu.Steal},
		{"guest", &cpu.Guest},
		{"guest_nice", &cpu.GuestNice},
	}

	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to scan /proc/stat")
	}
	for i, valStr := range strings.Fields(scanner.Text())[1:] {
		val, err := strconv.ParseUint(valStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to scan %s from /proc/stat", cpuStats[i].name)
		}
		*cpuStats[i].ptr = val
		cpu.Total += val
	}

	// Since cpustat[CPUTIME_USER] includes cpustat[CPUTIME_GUEST], substruct the duplicated values from total.
	// https://github.com/torvalds/linux/blob/4ec9f7a18/kernel/sched/cputime.c#L151-L158
	cpu.Total -= cpu.Guest
	// cpustat[CPUTIME_NICE] includes cpustat[CPUTIME_GUEST_NICE]
	cpu.Total -= cpu.GuestNice

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == 'c' && line[1] == 'p' && line[2] == 'u' && unicode.IsDigit(rune(line[3])) {
			cpu.CPUCount++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for /proc/stat: %s", err)
	}

	return &cpu, nil
}
