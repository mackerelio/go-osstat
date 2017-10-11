// +build !windows

package loadavg

import (
	"errors"
)

// #include <stdlib.h>
import "C"

func get() (*Loadavg, error) {
	var loadavgs [3]C.double
	ret := C.getloadavg(&loadavgs[0], 3)
	if ret != 3 {
		return nil, errors.New("failed to get loadavg")
	}
	return &Loadavg{
		Loadavg1:  float64(loadavgs[0]),
		Loadavg5:  float64(loadavgs[1]),
		Loadavg15: float64(loadavgs[2]),
	}, nil
}
