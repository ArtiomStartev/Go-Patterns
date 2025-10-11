package manualcancelation

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Launch a goroutine that will listen to the context cancellation
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 1 canceled: ", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 1 is working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Launch another goroutine that simulates some work
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 2 canceled: ", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 2 is working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Simulate some work in the main function
	fmt.Println("Main function is working...")
	time.Sleep(2 * time.Second)

	// Cancel the context, which will signal all goroutines to stop
	fmt.Println("Canceling context...")
	cancel()

	// Give goroutines time to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function done.")
}
