//go:build windows
// +build windows

package loadavg

import (
	"testing"
)

func TestGetLoadavg(t *testing.T) {
	loadavg, err := Get()
	if err == nil {
		t.Errorf("error should occur for Windows")
	}
	if loadavg != nil {
		t.Errorf("loadavg should be nil")
	}
}
