// +build freebsd

package memory

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetMemory(t *testing.T) {
	memory, err := Get()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	if memory.Used <= 0 || memory.Total <= 0 {
		t.Errorf("invalid memory value: %+v", memory)
	}
	t.Logf("memory value: %+v", memory)
}

func TestCollectMemoryStats(t *testing.T) {
	got, err := collectMemoryStats(strings.NewReader(
		`last pid:  1839;  load averages:  0.37,  0.35,  0.32  up 0+00:25:05    15:11:21
19 processes:  2 running, 17 sleeping

Mem: 17M Active, 644M Inact, 159M Wired, 61M Buf, 146M Free
Swap: 1024M Total, 1024M Free


  PID USERNAME    THR PRI NICE   SIZE    RES STATE    TIME    WCPU COMMAND
  739 root          6  52    0 26804K  3544K sigwai   0:00   0.00% VBoxService
`))
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	megaByte := 1024 * 1024
	expected := &Stats{
		Total:     uint64(0),
		Used:      uint64(0),
		Buffers:   uint64(61 * megaByte),
		Free:      uint64(146 * megaByte),
		Active:    uint64(17 * megaByte),
		Inactive:  uint64(644 * megaByte),
		Wired:     uint64(159 * megaByte),
		SwapTotal: uint64(1024 * megaByte),
		SwapFree:  uint64(1024 * megaByte),
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("invalid memory value: %+v (expected: %+v)", got, expected)
	}
}
