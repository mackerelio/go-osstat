// +build windows

package main

var generators []Generators

func init() {
	generators = []Generator{
		&cpuGenerator{},
		&memoryGenerator{},
		&networkGenerator{},
	}
}
