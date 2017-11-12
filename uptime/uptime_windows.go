// +build windows

package uptime

import (
	"syscall"
	"time"
)

var getTickCount = syscall.NewLazyDLL("kernel32.dll").NewProc("GetTickCount64")

func Get() (time.Duration, error) {
	ret, _, err := getTickCount.Call()
	if err != nil {
		return time.Duration(0), err
	}
	return time.Duration(ret) * time.Millisecond, nil
}
