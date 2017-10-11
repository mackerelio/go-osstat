// +build !windows

package loadavg

import (
	"testing"
)

func Test_GetLoadavg(t *testing.T) {
	loadavg, err := Get()
	if err != nil {
		t.Errorf("error should be nil but got: %v", err)
	}
	if loadavg.Loadavg1 < 0 || loadavg.Loadavg5 < 0 || loadavg.Loadavg15 < 0 {
		t.Errorf("invalid loadavg value: %v", loadavg)
	}
}
