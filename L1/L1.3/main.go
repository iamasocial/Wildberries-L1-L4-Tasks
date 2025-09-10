package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	count := flag.Int("workers", 5, "number of workers")
	flag.Parse()

	channel := make(chan int)

	var wg sync.WaitGroup
	wg.Add(*count)
	for i := 1; i <= *count; i++ {
		go worker(i, channel, &wg)
	}

	for i := 1; i <= 100; i++ {
		channel <- i
		time.Sleep(100 * time.Millisecond)
	}

	close(channel)
	wg.Wait()
}

func worker(id int, channel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range channel {
		fmt.Printf("Worker %d got %d\n", id, num)
	}
}
