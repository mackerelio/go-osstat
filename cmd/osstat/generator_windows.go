//go:build windows
// +build windows

package main

var generators []generator

func init() {
	generators = []generator{
		&uptimeGenerator{},
		&memoryGenerator{},
	}
}
