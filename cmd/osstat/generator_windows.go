// +build windows

package main

var generators []generator

func init() {
	generators = []generator{
		&memoryGenerator{},
	}
}
