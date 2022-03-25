//go:build linux
// +build linux

package disk

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetDisk(t *testing.T) {
	disks, err := Get()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	t.Logf("disks value: %+v", disks)
}

func TestCollectDiskStats(t *testing.T) {
	got, err := collectDiskStats(strings.NewReader(
		` 202       1 xvda1 750193 3037 28116978 368712 16600606 7233846 424712632 23987908 0 2355636 24345740
 202       2 xvda2 1641 9310 87552 1252 6365 3717 80664 24192 0 15040 25428
   7       0 loop0 0 0 0 0 0 0 0 0 0 0 0
   7       1 loop1 0 0 0 0 0 0 0 0 0 0 0
 253       0 dm-0 46095806 0 549095028 2243928 7192424 0 305024576 12521088 0 2728444 14782668
 253     628 dm-628 3198 0 75410 1360 30802835 0 3942653176 1334317408 0 70948 1358596768
 253       2 dm-2 2022 0 42250 488 30822403 0 3942809696 1364721232 0 93348 1382989868
`))
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	expected := []Stats{
		{"xvda1", 750193, 16600606},
		{"xvda2", 1641, 6365},
		{"loop0", 0, 0},
		{"loop1", 0, 0},
		{"dm-0", 46095806, 7192424},
		{"dm-628", 3198, 30802835},
		{"dm-2", 2022, 30822403},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("invalid disk value: %+v (expected: %+v)", got, expected)
	}
}
