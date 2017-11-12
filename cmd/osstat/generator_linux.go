package main

var generators []generator

func init() {
	generators = []generator{
		&loadavgGenerator{},
		&uptimeGenerator{},
		&cpuGenerator{},
		&memoryGenerator{},
		&diskGenerator{},
		&networkGenerator{},
	}
}
