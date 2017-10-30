// +build darwin

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
		`Mach Virtual Memory Statistics: (page size of 4096 bytes)
Pages free:                               72827.
Pages active:                           2154445.
Pages inactive:                         1511468.
Pages speculative:                         8107.
Pages throttled:                              0.
Pages wired down:                        446975.
Pages purgeable:                         383371.
"Translation faults":                  97589077.
Pages copy-on-write:                    3305869.
Pages zero filled:                     50848672.
Pages reactivated:                         1999.
Pages purged:                           2496610.
File-backed pages:                       677870.
Anonymous pages:                        2996150.
Pages stored in compressor:                   0.
Pages occupied by compressor:                 0.
Decompressions:                               0.
Compressions:                                 0.
Pageins:                                6333901.
Pageouts:                                   353.
Swapins:                                      0.
Swapouts:                                     0.
`))
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	pageSize := 4096
	expected := &Stats{
		Total:    uint64((446975 + 2154445 + 1511468 + 8107 + 72827) * pageSize),
		Used:     uint64((446975 + 2154445 + 1511468 + 8107 - (383371 + 677870)) * pageSize),
		Cached:   uint64((383371 + 677870) * pageSize),
		Free:     uint64(72827 * pageSize),
		Active:   uint64(2154445 * pageSize),
		Inactive: uint64(1511468 * pageSize),
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("invalid memory value: %+v (expected: %+v)", got, expected)
	}
}

func TestCollectSwapStats(t *testing.T) {
	got, err := collectSwapStats([]byte(
		"\x00\x00\x00\x40\x01\x00\x00\x00\x00\x00\x3c\x56\x00\x00\x00\x00\x00\x00\xc4\xe9\x00\x00\x00\x00\x00\x10\x00\x00\x01\x00\x00\x00",
	))
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	expected := &swapUsage{
		Total:     0x0000000140000000,
		Avail:     0x00000000563c0000,
		Used:      0x00000000e9c40000,
		Pagesize:  4096,
		Encrypted: true,
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("invalid memory value: %+v (expected: %+v)", got, expected)
	}
}
