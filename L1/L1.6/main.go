package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go stopByCondition(&wg)
	wg.Wait()
	fmt.Println()

	wg.Add(1)
	go stopByChannel(&wg)
	wg.Wait()
	fmt.Println()

	wg.Add(1)
	go stopByContext(&wg)
	wg.Wait()
	fmt.Println()

	wg.Add(1)
	go stopByTimer(&wg)
	wg.Wait()
	fmt.Println()

	wg.Add(1)
	go stopByGoexit(&wg)
	wg.Wait()
}

func stopByCondition(wg *sync.WaitGroup) {
	defer wg.Done()
	stop := false

	go func() {
		for !stop {
			fmt.Printf("Running, stop = false (condition)\n")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Printf("Goroutine stopping, stop = true (condition)\n")
	}()

	time.Sleep(3 * time.Second)
	stop = true
	time.Sleep(1 * time.Second)
}

func stopByChannel(wg *sync.WaitGroup) {
	defer wg.Done()

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Printf("Goroutine stopping, channel is closed (channnel)\n")
				return
			default:
				fmt.Printf("Running, channel is open (channel)\n")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
}

func stopByContext(wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Goroutine stopping, cancel() (context)\n")
				return
			default:
				fmt.Printf("Running (context)\n")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func stopByTimer(wg *sync.WaitGroup) {
	defer wg.Done()

	timer := time.NewTimer(3 * time.Second)

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Printf("Goroutine stopping, timer expired (timer)\n")
				return
			default:
				fmt.Printf("Running (timer)\n")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(4 * time.Second)
}

func stopByGoexit(wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Running (goexit)\n")
			time.Sleep(500 * time.Millisecond)
			if i == 5 {
				fmt.Printf("Goroutine stopping, calling runtime.Goexit()\n")
				runtime.Goexit()
			}
		}
	}()

	time.Sleep(4 * time.Second)
}
