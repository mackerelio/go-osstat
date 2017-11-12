// +build darwin freebsd netbsd

package uptime

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

func get() (time.Duration, error) {
	out, err := unix.SysctlRaw("kern.boottime")
	if err != nil {
		return time.Duration(0), err
	}
	var timeval syscall.Timeval
	if len(out) != int(unsafe.Sizeof(timeval)) {
		return time.Duration(0), fmt.Errorf("unexpected output of sysctl kern.boottime: %v (len: %d)", out, len(out))
	}
	timeval = *(*syscall.Timeval)(unsafe.Pointer(&out[0]))
	sec, nsec := timeval.Unix()
	return time.Now().Sub(time.Unix(sec, nsec)), nil
}
