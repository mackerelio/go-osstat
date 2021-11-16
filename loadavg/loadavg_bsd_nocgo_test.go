//go:build (darwin || freebsd || netbsd || openbsd) && !cgo
// +build darwin freebsd netbsd openbsd
// +build !cgo

package loadavg

import (
	"reflect"
	"testing"
)

func TestCollectLoadavgStats(t *testing.T) {
	got, err := collectLoadavgStats([]byte(
		"\xd6\x11\x00\x00\x92\x13\x00\x00\xfc\x12\x00\x00\x00\x00\x00\x00\x00\x08\x00\x00\x00\x00\x00\x00",
	))
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	expected := &Stats{Loadavg1: 2.2294921875, Loadavg5: 2.4462890625, Loadavg15: 2.373046875}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("invalid loadavg value: %+v (expected: %+v)", got, expected)
	}
}
