// +build windows

package loadavg

import (
	"errors"
)

func get() (*Loadavg, error) {
	return nil, errors.New("loadavg for Windows is not supported")
}
