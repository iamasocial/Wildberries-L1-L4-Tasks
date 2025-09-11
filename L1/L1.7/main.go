package main

import (
	"fmt"
	"sync"
)

type MySyncMap struct {
	m  map[int]string
	mu sync.Mutex
}

func New() *MySyncMap {
	return &MySyncMap{
		m:  make(map[int]string),
		mu: sync.Mutex{},
	}
}

func (sm *MySyncMap) Get(key int) (string, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	val, ok := sm.m[key]
	return val, ok
}

func (sm *MySyncMap) Set(key int, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.m[key] = value
}

func main() {
	syncMap := New()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				syncMap.Set(j, fmt.Sprintf("Goroutine %d - Value %d", i, j))
				fmt.Printf("Goroutine %d set key %d\n", i, j)
			}
		}(i)
	}

	wg.Wait()

	for k, v := range syncMap.m {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}

	fmt.Println("All goroutines write sequentially to the map")
	// fmt.Println()
	// fmt.Println("Now demonstrating concurrent writes to a standard map (not thread-safe)")

	// m := make(map[int]string)

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		for j := 0; j < 3; j++ {
	// 			m[j] = fmt.Sprintf("Goroutine %d - Value %d", i, j)
	// 			fmt.Printf("Goroutine %d set key %d\n", i, j)
	// 		}
	// 	}(i)
	// }

	// wg.Wait()

	// for k, v := range m {
	// 	fmt.Printf("Key: %d, Value: %s\n", k, v)
	// }
}
