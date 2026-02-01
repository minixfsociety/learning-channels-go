package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for n := range ch {
			fmt.Println("Worker: processing job", n)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)
	wg.Wait()
	fmt.Println("Main: All workers finished!")
}
