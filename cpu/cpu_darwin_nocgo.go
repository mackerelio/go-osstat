// +build darwin,!cgo

package cpu

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

// Get cpu statistics
func Get() (*CPU, error) {
	return nil, fmt.Errorf("cpu statistics for darwin is unavailable yet")
	cmd := exec.Command("iostat", "-n0", "-c2")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	cpu, err := collectCPUStats(out)
	if err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return cpu, nil
}

// CPU represents cpu statistics for darwin
type CPU struct {
	User, System, Idle, Total float64
}

func collectCPUStats(out io.Reader) (*CPU, error) {
	scanner := bufio.NewScanner(out)
	for i := 0; i < 4; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("failed to scan output of iostat")
		}
	}

	var cpu CPU
	line := scanner.Text()
	ret, err := fmt.Sscanf(line, "%f %f %f", &cpu.User, &cpu.System, &cpu.Idle)
	if err != nil || ret != 3 {
		return nil, fmt.Errorf("unexpected output of iostat")
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for iostat: %s", err)
	}

	cpu.Total = cpu.User + cpu.System + cpu.Idle
	return &cpu, nil
}
