package main

import (
	"io"
	"sync"
)

type Generator interface {
	Get()
	Error() error
	Print(out io.Writer)
}

func Run(args []string, out io.Writer) []error {
	var wg sync.WaitGroup
	generators := []Generator{
		&loadavgGenerator{},
		&cpuGenerator{},
		&memoryGenerator{},
		&networkGenerator{},
	}

	for _, generator := range generators {
		wg.Add(1)
		go func(generator Generator) {
			defer wg.Done()
			generator.Get()
		}(generator)
	}

	wg.Wait()

	var errs []error
	for _, generator := range generators {
		if err := generator.Error(); err != nil {
			errs = append(errs, err)
		} else {
			generator.Print(out)
		}
	}
	return errs
}
