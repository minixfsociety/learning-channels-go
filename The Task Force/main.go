package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		fmt.Printf("Worker %d: started job %d\n", id, n)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	fmt.Println("All jobs finished!")
}
