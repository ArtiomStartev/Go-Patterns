package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Create a context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Perform a search with the context
	res, err := Search(ctx, "random string")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Response: %s\n", res)
}

func Search(ctx context.Context, query string) (string, error) {
	// Channel to receive the response from the slow function
	resp := make(chan string)

	go func() {
		resp <- RandomSleepAndReturnAPI(query)
		close(resp)
	}()

	// Wait for either response or the context to be done
	for {
		select {
		case dst := <-resp:
			return dst, nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}

}

// RandomSleepAndReturnAPI simulates a slow API call by sleeping for a random duration
func RandomSleepAndReturnAPI(query string) string {
	randomDuration := time.Duration(rand.Intn(10))

	// Sleep for the random duration (up to 10 seconds)
	time.Sleep(randomDuration * time.Second)

	return fmt.Sprintf("It took us %v... Hope it was worth the wait! 🧭", randomDuration)
}
