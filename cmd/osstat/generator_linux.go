package main

var generators []generator

func init() {
	generators = []generator{
		&loadavgGenerator{},
		&cpuGenerator{},
		&memoryGenerator{},
		&diskGenerator{},
		&networkGenerator{},
	}
}
