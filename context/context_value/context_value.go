package main

import (
	"context"
	"log"
	"math/rand"
)

func main() {
	// Create a new context and store a value in it
	ctx := context.WithValue(context.Background(), "userId", rand.Intn(100))

	// Pass the context to a function that needs access to the value
	ProcessRequest(ctx)
}

// ProcessRequest simulates processing a request with access to context values
func ProcessRequest(ctx context.Context) {
	// Extract the value form the context using the same key
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		log.Println("userId not found in context")
		return
	}

	// Use the extracted value
	log.Printf("Processing request for user ID: %d\n", userId)
}
