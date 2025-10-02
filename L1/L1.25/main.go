package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	sleepTime := 2 * time.Second

	fmt.Printf("Before\tBusySleep:\t%v\n", time.Now())
	BusySleep(sleepTime)
	fmt.Printf("After\tBusySleep:\t%v\n\n", time.Now())

	fmt.Printf("Before\tChanSleep:\t%v\n", time.Now())
	ChanSleep(sleepTime)
	fmt.Printf("After\tChanSleep:\t%v\n\n", time.Now())

	fmt.Printf("Before\tTimerSleep:\t%v\n", time.Now())
	TimerSleep(sleepTime)
	fmt.Printf("After\tTimerSleep:\t%v\n\n", time.Now())

	fmt.Printf("Before\tAfterSleep:\t%v\n", time.Now())
	AfterSleep(sleepTime)
	fmt.Printf("After\tAfterSleep:\t%v\n\n", time.Now())

	fmt.Printf("Before\tContextSleep:\t%v\n", time.Now())
	ContextSleep(sleepTime)
	fmt.Printf("After\tContextSleep:\t%v\n", time.Now())

}

func BusySleep(t time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) > t {
			return
		}
	}
}

func ChanSleep(t time.Duration) {
	ch := make(chan struct{})
	time.AfterFunc(t, func() {
		close(ch)
	})
	<-ch
}

func AfterSleep(t time.Duration) {
	<-time.After(t)
}

func TimerSleep(t time.Duration) {
	timer := time.NewTimer(t)
	<-timer.C
}

func ContextSleep(t time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	<-ctx.Done()
}
