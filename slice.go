package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var xi []*int

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			xi = append(xi, &ii)
		}(i)
	}
	wg.Wait()

	fmt.Println(len(xi))
}
