package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int)
	double := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)

	go inputWorker(numbers, ch, &wg)
	go doubleWorker(ch, double, &wg)
	go outputWorker(double, &wg)

	wg.Wait()
	fmt.Printf("All goroutines done")
}

func inputWorker(numbers []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	for _, num := range numbers {
		ch <- num
	}
}

func doubleWorker(orig <-chan int, double chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(double)

	for num := range orig {
		double <- num * 2
	}
}

func outputWorker(double <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range double {
		fmt.Printf("%d\n", num)
	}
}
