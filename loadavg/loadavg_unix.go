// +build !windows

package loadavg

import (
	"errors"
	"unsafe"
)

// #include <stdlib.h>
import "C"

func get() (*Loadavg, error) {
	var loadavgs [3]float64
	ret := C.getloadavg((*C.double)(unsafe.Pointer(&loadavgs)), 3)
	if ret != 3 {
		return nil, errors.New("failed to get loadavg")
	}
	return &Loadavg{
		Loadavg1:  loadavgs[0],
		Loadavg5:  loadavgs[1],
		Loadavg15: loadavgs[2],
	}, nil
}
