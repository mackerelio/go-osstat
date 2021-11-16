//go:build darwin && !cgo
// +build darwin,!cgo

package main

var generators []generator

func init() {
	generators = []generator{
		&loadavgGenerator{},
		&uptimeGenerator{},
		&memoryGenerator{},
		&networkGenerator{},
	}
}
