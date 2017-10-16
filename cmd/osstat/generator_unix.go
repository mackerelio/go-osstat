// +build !windows,!linux,!darwin darwin,cgo

package main

var generators []generator

func init() {
	generators = []generator{
		&loadavgGenerator{},
		&cpuGenerator{},
		&memoryGenerator{},
		&networkGenerator{},
	}
}
