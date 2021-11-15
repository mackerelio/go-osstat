//go:build linux
// +build linux

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
	testCases := []struct {
		name   string
		input  string
		expect *Stats
	}{
		{
			name: "Disable MemAvailable",
			input: `MemTotal:        1929620 kB
MemFree:          113720 kB
Buffers:           81744 kB
Cached:           435712 kB
SwapCached:          504 kB
Active:           817412 kB
Inactive:         754140 kB
Active(anon):     647484 kB
Inactive(anon):   570160 kB
Active(file):     169928 kB
Inactive(file):   183980 kB
Unevictable:         124 kB
Mlocked:             124 kB
HighTotal:       1047928 kB
HighFree:          18692 kB
LowTotal:         881692 kB
LowFree:           95028 kB
SwapTotal:       1959932 kB
SwapFree:        1957500 kB
Dirty:               352 kB
Writeback:             0 kB
AnonPages:       1053804 kB
Mapped:           151408 kB
Shmem:            163548 kB
Slab:             202768 kB
SReclaimable:     177128 kB
SUnreclaim:        25640 kB
KernelStack:        4624 kB
PageTables:        15944 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:     2924740 kB
Committed_AS:    7238800 kB
VmallocTotal:     122880 kB
VmallocUsed:       16344 kB
VmallocChunk:     102740 kB
HardwareCorrupted:     0 kB
AnonHugePages:    145408 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
DirectMap4k:       24568 kB
DirectMap2M:      888832 kB
`,
			expect: &Stats{
				Total:               uint64(1929620 * 1024),
				Used:                uint64(1298444 * 1024),
				Buffers:             uint64(81744 * 1024),
				Cached:              uint64(435712 * 1024),
				Free:                uint64(113720 * 1024),
				Active:              uint64(817412 * 1024),
				Inactive:            uint64(754140 * 1024),
				SwapTotal:           uint64(1959932 * 1024),
				SwapUsed:            uint64(2432 * 1024),
				SwapCached:          uint64(504 * 1024),
				SwapFree:            uint64(1957500 * 1024),
				Mapped:              uint64(151408 * 1024),
				Shmem:               uint64(163548 * 1024),
				Slab:                uint64(202768 * 1024),
				PageTables:          uint64(15944 * 1024),
				Committed:           uint64(7238800 * 1024),
				VmallocUsed:         uint64(16344 * 1024),
				MemAvailableEnabled: false,
			},
		},
		{
			name: "Enable MemAvailable",
			input: `MemTotal:        1929620 kB
MemFree:          113720 kB
MemAvailable:     533132 kB
Buffers:           81744 kB
Cached:           435712 kB
SwapCached:          504 kB
Active:           817412 kB
Inactive:         754140 kB
Active(anon):     647484 kB
Inactive(anon):   570160 kB
Active(file):     169928 kB
Inactive(file):   183980 kB
Unevictable:         124 kB
Mlocked:             124 kB
HighTotal:       1047928 kB
HighFree:          18692 kB
LowTotal:         881692 kB
LowFree:           95028 kB
SwapTotal:       1959932 kB
SwapFree:        1957500 kB
Dirty:               352 kB
Writeback:             0 kB
AnonPages:       1053804 kB
Mapped:           151408 kB
Shmem:            163548 kB
Slab:             202768 kB
SReclaimable:     177128 kB
SUnreclaim:        25640 kB
KernelStack:        4624 kB
PageTables:        15944 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:     2924740 kB
Committed_AS:    7238800 kB
VmallocTotal:     122880 kB
VmallocUsed:       16344 kB
VmallocChunk:     102740 kB
HardwareCorrupted:     0 kB
AnonHugePages:    145408 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
DirectMap4k:       24568 kB
DirectMap2M:      888832 kB
`,
			expect: &Stats{
				Total:               uint64(1929620 * 1024),
				Used:                uint64(1396488 * 1024),
				Buffers:             uint64(81744 * 1024),
				Cached:              uint64(435712 * 1024),
				Free:                uint64(113720 * 1024),
				Available:           uint64(533132 * 1024),
				Active:              uint64(817412 * 1024),
				Inactive:            uint64(754140 * 1024),
				SwapTotal:           uint64(1959932 * 1024),
				SwapUsed:            uint64(2432 * 1024),
				SwapCached:          uint64(504 * 1024),
				SwapFree:            uint64(1957500 * 1024),
				Mapped:              uint64(151408 * 1024),
				Shmem:               uint64(163548 * 1024),
				Slab:                uint64(202768 * 1024),
				PageTables:          uint64(15944 * 1024),
				Committed:           uint64(7238800 * 1024),
				VmallocUsed:         uint64(16344 * 1024),
				MemAvailableEnabled: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := collectMemoryStats(strings.NewReader(tc.input))
			if err != nil {
				t.Fatalf("error should be nil but got: %v", err)
			}
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("%s: invalid memory value: %+v (expected: %+v)", tc.name, got, tc.expect)
			}
		})
	}
}
