package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	count := flag.Int("workers", 5, "number of workers")
	flag.Parse()

	channel := make(chan int)

	var wg sync.WaitGroup
	wg.Add(*count)
	for i := 1; i <= *count; i++ {
		go worker(i, channel, ctx, &wg)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Printf("\nReceived interrupt, shutting down\n")
		cancel()
	}()

Loop:
	for i := 1; i <= 100; i++ {
		select {
		case channel <- i:
			time.Sleep(500 * time.Millisecond)
		case <-ctx.Done():
			fmt.Printf("Main context cancelled, stopping sending\n")
			break Loop
		}
	}

	close(channel)
	wg.Wait()

	fmt.Println("All workers stopped, exiting program")
}

func worker(id int, channel <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num, ok := <-channel:
			if !ok {
				fmt.Printf("Channel closed, worker %d exiting\n", id)
				return
			}
			fmt.Printf("Worker %d got %d\n", id, num)
		case <-ctx.Done():
			fmt.Printf("Context cancelled, worker %d exiting\n", id)
			return
		}
	}
}
