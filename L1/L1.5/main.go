package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout duration")
	flag.Parse()

	timer := time.NewTimer(*timeout)

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	go writer(ch, timer, &wg)
	go reader(ch, &wg)

	wg.Wait()
	fmt.Printf("Timeout reached, stopping program\n")

}

func writer(ch chan<- int, timer *time.Timer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; ; i++ {
		select {
		case <-timer.C:
			close(ch)
			fmt.Printf("Timer expired, writer exiting\n")
			return
		case ch <- i:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func reader(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			fmt.Printf("Channel closed, reader exiting\n")
			return
		}
		fmt.Printf("Reader received: %d\n", num)
	}
}
