package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	out := bufio.NewWriter(new(bytes.Buffer))
	errs := run(nil, out)
	if errs != nil {
		t.Errorf("error occured: %v", errs)
	}
}
