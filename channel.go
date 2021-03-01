package main

import (
	"fmt"
	"sync"
)

const noOfWorkers = 4

func main() {
	var xi []int
	var wg sync.WaitGroup

	queue := make(chan int, noOfWorkers)
	defer close(queue)

	// Create our data and send it into the queue.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(ii int) {
			queue <- ii
		}(i)
	}

	go func() {
		for t := range queue {
			xi = append(xi, t)
			wg.Done()
		}
	}()
	wg.Wait()

	fmt.Println(len(xi))
}
