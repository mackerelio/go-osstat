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
	var loadavgs Loadavg
	ret, err := fmt.Fscanf(out, "%f %f %f", &loadavgs.Loadavg1, &loadavgs.Loadavg5, &loadavgs.Loadavg15)
	if err != nil || ret != 3 {
		return nil, fmt.Errorf("unexpected format of /proc/loadavg")
	}
	return &loadavgs, nil
}
