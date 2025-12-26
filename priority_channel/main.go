package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	jobs := make(chan int)

	// Worker goroutine
	go func() {
		for {
			// 1. High-priority channel (non-blocking)
			select {
			case <-stop:
				fmt.Println("[worker] STOP signal received")
				return
			default:
			}

			// 2. Normal work (blocking)
			select {
			case job := <-jobs:
				fmt.Println("[worker] processing job:", job)
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	// Send some jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
		time.Sleep(200 * time.Millisecond)
	}

	// Send stop signal
	fmt.Println("[main] sending STOP")
	close(stop)

	// Give worker time to exit
	time.Sleep(time.Second)
	fmt.Println("[main] done")
}
