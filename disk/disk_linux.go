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

// Get disk I/O statistics
func Get() ([]DiskStats, error) {
	file, err := os.Open("/proc/diskstats")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return collectDiskStats(file)
}

// DiskStats represents disk I/O statistics for linux
type DiskStats struct {
	Name                            string
	ReadsCompleted, WritesCompleted uint64
}

func collectDiskStats(out io.Reader) ([]DiskStats, error) {
	scanner := bufio.NewScanner(out)
	var diskStats []DiskStats
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 14 {
			continue
		}
		// Reference: Documentation/iostats.txt in the source of Linux
		name := fields[2]
		readsCompleted, err := strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse reads completed of %s", name)
		}
		writesCompleted, err := strconv.ParseUint(fields[7], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse writes completed of %s", name)
		}
		diskStats = append(diskStats, DiskStats{
			Name:            name,
			ReadsCompleted:  readsCompleted,
			WritesCompleted: writesCompleted,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for /proc/net/dev: %s", err)
	}
	return diskStats, nil
}
