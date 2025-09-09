package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}
	wg.Add(len(numbers))
	for _, number := range numbers {
		go func() {
			defer wg.Done()
			res := square(number)
			fmt.Printf("%d^2 = %d\n", number, res)
		}()
	}
	wg.Wait()
}

func square(n int) int {
	return n * n
}
