// +build freebsd

package memory

import (
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

func TestCollectSwapStats(t *testing.T) {
	total, used, err := collectSwapStats(strings.NewReader(
		`Device          1K-blocks     Used    Avail Capacity
/dev/gpt/swapfs   1048576    16056  1032520     2%
`))
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}

	totalExpected, usedExpected := uint64(1048576)*1024, uint64(16056)*1024
	if total != totalExpected {
		t.Errorf("invalid swap total: %+v (expected: %+v)", total, totalExpected)
	}
	if used != usedExpected {
		t.Errorf("invalid swap used: %+v (expected: %+v)", used, usedExpected)
	}
}
