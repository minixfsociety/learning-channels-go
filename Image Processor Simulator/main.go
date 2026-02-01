package main

import (
	"fmt"
	"sync"
	"time"
)

func processor(id int, task <-chan int, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range task {
		time.Sleep(1 * time.Second)
		result <- fmt.Sprintf("Worker %d finished task %d", id, n)
	}
}

func main() {
	task := make(chan int)
	result := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go processor(i, task, result, &wg)
	}
	go func() {
		for i := 0; i < 5; i++ {
			task <- i
		}
		close(task)
	}()
	go func() {
		wg.Wait()
		close(result)
	}()
	for res := range result {
		fmt.Println(res)
	}
}
