//go:build linux
// +build linux

package disk

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Get disk I/O statistics.
func Get() ([]Stats, error) {
	// Reference: Documentation/iostats.txt in the source of Linux
	file, err := os.Open("/proc/diskstats")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return collectDiskStats(file)
}

// Stats represents disk I/O statistics for linux.
type Stats struct {
	Name            string // device name; like "hda"
	ReadsCompleted  uint64 // total number of reads completed successfully
	WritesCompleted uint64 // total number of writes completed successfully
}

func collectDiskStats(out io.Reader) ([]Stats, error) {
	scanner := bufio.NewScanner(out)
	var diskStats []Stats
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 14 {
			continue
		}
		name := fields[2]
		readsCompleted, err := strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse reads completed of %s", name)
		}
		writesCompleted, err := strconv.ParseUint(fields[7], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse writes completed of %s", name)
		}
		diskStats = append(diskStats, Stats{
			Name:            name,
			ReadsCompleted:  readsCompleted,
			WritesCompleted: writesCompleted,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for /proc/diskstats: %s", err)
	}
	return diskStats, nil
}
