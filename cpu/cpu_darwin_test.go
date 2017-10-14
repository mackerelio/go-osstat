// +build darwin

package cpu

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetCpu(t *testing.T) {
	cpu, err := Get()
	if err != nil {
		t.Errorf("error should be nil but got: %v", err)
	}
	if cpu.User <= 0 || cpu.System <= 0 {
		t.Errorf("invalid cpu value: %+v", cpu)
	}
}

func TestCollectCpuStats(t *testing.T) {
	got, err := collectCpuStats(strings.NewReader(
		`      cpu    load average
 us sy id   1m   5m   15m
 12  6 83  2.28 2.18 2.19
 10  7 82  2.28 2.18 2.19
`))
	if err != nil {
		t.Errorf("error should be nil but got: %v", err)
	}
	expected := &Cpu{User: 10.0, System: 7.0, Idle: 82.0}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("invalid cpu value: %+v (expected: %+v)", got, expected)
	}
}
