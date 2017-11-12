// +build linux

package uptime

import (
	"syscall"
	"time"
)

func get() (time.Duration, error) {
	var info syscall.Sysinfo_t
	if err := syscall.Sysinfo(&info); err != nil {
		return time.Duration(0), err
	}
	return time.Duration(info.Uptime) * time.Second, nil
}
