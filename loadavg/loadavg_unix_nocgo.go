// +build linux,!cgo
// +build !darwin

package loadavg

import (
	"fmt"
	"io"
	"os"
)

func get() (*Loadavg, error) {
	file, err := os.Open("/proc/loadavg")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return collectLoadavgStats(file)
}

func collectLoadavgStats(out io.Reader) (*Loadavg, error) {
	var loadavg Loadavg
	ret, err := fmt.Fscanf(out, "%f %f %f", &loadavg.Loadavg1, &loadavg.Loadavg5, &loadavg.Loadavg15)
	if err != nil || ret != 3 {
		return nil, fmt.Errorf("unexpected format of /proc/loadavg")
	}
	return &loadavg, nil
}
