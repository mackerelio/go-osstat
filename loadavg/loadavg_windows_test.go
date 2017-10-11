// +build windows

package loadavg

import (
	"testing"
)

func Test_GetLoadavg(t *testing.T) {
	loadavg, err := Get()
	if err == nil {
		t.Errorf("error should occur for Windows")
	}
}
