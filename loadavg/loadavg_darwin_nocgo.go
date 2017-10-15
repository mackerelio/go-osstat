// +build darwin,!cgo

package loadavg

import (
	"fmt"
	"io"
	"os/exec"
)

func get() (*Loadavg, error) {
	cmd := exec.Command("sysctl", "-n", "vm.loadavg")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	loadavg, err := collectLoadavgStats(out)
	if err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return loadavg, nil
}

func collectLoadavgStats(out io.Reader) (*Loadavg, error) {
	var loadavg Loadavg
	ret, err := fmt.Fscanf(out, "{ %f %f %f }", &loadavg.Loadavg1, &loadavg.Loadavg5, &loadavg.Loadavg15)
	if err != nil || ret != 3 {
		return nil, fmt.Errorf("unexpected output of sysctl -n vm.loadavg")
	}
	return &loadavg, nil
}
