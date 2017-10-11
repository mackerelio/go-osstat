// +build windows

package loadavg

import (
	"errors"
)

func get() (Loadavg, error) {
	return Loadavg{}, errors.New("loadavg for Windows is not supported")
}
