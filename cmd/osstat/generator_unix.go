// +build !windows

package main

var generators []Generator

func init() {
	generators = []Generator{
		&loadavgGenerator{},
		&cpuGenerator{},
		&memoryGenerator{},
		&networkGenerator{},
	}
}
