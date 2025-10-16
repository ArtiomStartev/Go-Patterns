package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"
)

func fetchTodosWithWeightedSemaphore(ctx context.Context, ids []int) []TodoItem {
	var wg sync.WaitGroup
	todos := make([]TodoItem, len(ids))
	sem := semaphore.NewWeighted(int64(runtime.NumCPU()))

	for i, id := range ids {
		id := id

		sem.Acquire(ctx, 1) // // acquire slot
		wg.Add(1)

		go func() {
			defer func() {
				sem.Release(1) // release slot
				wg.Done()
			}()

			todo, err := fetchTodo(ctx, id)
			if err != nil {
				fmt.Println("Error fetching todo: ", err)
				return
			}

			todos[i] = todo
		}()
	}

	wg.Wait()

	return todos
}
