package main

import (
	"context"
	"fmt"
	"time"
)

const (
	requestTimeout = 10 * time.Second
)

type todoFetcher func(context.Context, []int) []TodoItem

func main() {
	todoIDs := buildTodoIDs()

	examples := []struct {
		label string
		fetch todoFetcher
	}{
		{"Buffered channel semaphore", fetchTodosWithChannel},
		{"golang.org/x/sync/semaphore", fetchTodosWithWeightedSemaphore},
		{"golang.org/x/sync/errgroup", fetchTodosWithErrGroup},
	}

	for _, example := range examples {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		fmt.Printf("=== %s ===\n", example.label)

		todos := example.fetch(ctx, todoIDs)
		cancel()

		printTodos(todos)
		fmt.Println()
	}
}

func buildTodoIDs() []int {
	ids := make([]int, 200)
	for i := range ids {
		ids[i] = i + 1
	}
	return ids
}
