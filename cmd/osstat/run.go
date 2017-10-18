package main

import (
	"fmt"
	"io"
	"sync"
)

type generator interface {
	Get()
	Error() error
	Print(out chan<- value)
}

type value struct {
	name  string
	value interface{}
	unit  string
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

	c := make(chan value)
	done := make(chan struct{})
	defer close(done)

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Fprintf(out, "%-25s %-14v %s\n", v.name, v.value, v.unit)
			case <-done:
				close(c)
				return
			}
		}
	}()

	var errs []error
	for _, gen := range generators {
		if err := gen.Error(); err != nil {
			errs = append(errs, err)
		} else {
			gen.Print(c)
		}
	}
	done <- struct{}{}

	return errs
}
