package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

var justString string

func main() {
	doCopy := flag.Bool("copy", false, "copy substring to avoid leak (safe mode)")
	iter := flag.Int("iter", 10, "how many iterations")
	sizeMB := flag.Int("mb", 10, "size of created string in megabytes")
	delay := flag.Duration("delay", 500*time.Millisecond, "ms delay between iterations")
	flag.Parse()

	size := *sizeMB << 20

	fmt.Printf("Leak demo: copy=%v, iterations=%d, size=%d MB\n", *doCopy, *iter, *sizeMB)

	var m runtime.MemStats

	for i := 0; i < *iter; i++ {

		if *doCopy {
			someFuncSafe(size)
		} else {
			someFuncLeak(size)
		}

		runtime.GC()
		runtime.ReadMemStats(&m)

		fmt.Printf("iter %2d: HeapAlloc = %12d bytes (justString len=%d)\n", i+1, m.HeapAlloc, len(justString))

		time.Sleep(*delay)
	}
}

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFuncLeak(size int) {
	v := createHugeString(size)
	justString = v[:100]
}

func someFuncSafe(size int) {
	v := createHugeString(size)
	justString = string([]byte(v[:100]))
}
