package uptime

import (
	"testing"
)

func TestGetUptime(t *testing.T) {
	uptime, err := Get()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	if uptime.Seconds() <= 0 {
		t.Errorf("invalid uptime value: %v", uptime)
	}
	t.Logf("uptime value: %+v", uptime)
}
