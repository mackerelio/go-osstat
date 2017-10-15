package main

import (
	"io"
	"sync"
)

type generator interface {
	Get()
	Error() error
	Print(out io.Writer)
}

func run(args []string, out io.Writer) []error {
	var wg sync.WaitGroup

	for _, gen := range generators {
		wg.Add(1)
		go func(gen generator) {
			defer wg.Done()
			gen.Get()
		}(gen)
	}

	wg.Wait()

	var errs []error
	for _, gen := range generators {
		if err := gen.Error(); err != nil {
			errs = append(errs, err)
		} else {
			gen.Print(out)
		}
	}
	return errs
}
