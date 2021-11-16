//go:build linux
// +build linux

package cpu

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetCPU(t *testing.T) {
	cpu, err := Get()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	if cpu.User <= 0 || cpu.System <= 0 || cpu.Total <= 0 || cpu.StatCount < 4 {
		t.Errorf("invalid cpu value: %+v", cpu)
	}
	t.Logf("cpu value: %+v", cpu)
}

func TestCollectCPUStats(t *testing.T) {
	got, err := collectCPUStats(strings.NewReader(
		`cpu  1415984 38486 429451 2500643 10585 157 2372 0 0 0
cpu0 708614 19410 217184 2188812 9733 144 808 0 0 0
cpu1 707370 19076 212266 311830 851 12 1564 0 0 0
intr 40269386 11401108 2407 0 0 0 0 0 0 1 2601 0 0 914 0 0 0 360 0 0 21183 0 54 0 16365 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 839980 2127556 1919962 429 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0

ctxt 151685704
btime 1507943277
processes 28087
procs_running 8
procs_blocked 0
softirq 10624366 42 5280893 11772 27757 826862 2 24721 2326791 28519 2097007
`))
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	expected := &Stats{
		User:      1415984,
		Nice:      38486,
		System:    429451,
		Idle:      2500643,
		Iowait:    10585,
		Irq:       157,
		Softirq:   2372,
		Steal:     0,
		Guest:     0,
		GuestNice: 0,
		Total:     4397678,
		CPUCount:  2,
		StatCount: 10,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("invalid cpu value: %+v (expected: %+v)", got, expected)
	}
}
