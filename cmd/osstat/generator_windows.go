// +build windows

package main

var generators []generator

func init() {
	generators = []generator{
		&cpuGenerator{},
		&memoryGenerator{},
		&networkGenerator{},
	}
}
