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
		time.Sleep(3 * time.Second)
		stop = true
	}()

	for !stop {
		fmt.Printf("Running, stop = false (condition)\n")
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("Goroutine stopping, stop = true (condition)\n")
}

func stopByChannel(wg *sync.WaitGroup) {
	defer wg.Done()

	ch := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()

	for {
		select {
		case <-ch:
			fmt.Printf("Goroutine stopping, channel closed (channel)\n")
			return
		default:
			fmt.Printf("Running, channel open (channel)\n")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func stopByContext(wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Goroutine stopping, context cancelled (context)\n")
			return
		default:
			fmt.Printf("Running, context active (context)\n")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func stopByTimer(wg *sync.WaitGroup) {
	defer wg.Done()

	timer := time.NewTimer(3 * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Printf("Goroutine stopping, timer expired (timer)\n")
			return
		default:
			fmt.Printf("Running, timer active (timer)\n")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func stopByGoexit(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Printf("Running (goexit)\n")
		time.Sleep(500 * time.Millisecond)
		if i == 6 {
			fmt.Printf("Goroutine stopping (goexit)\n")
			runtime.Goexit()
		}
	}
}
