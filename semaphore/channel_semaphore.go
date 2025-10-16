package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func fetchTodosWithChannel(ctx context.Context, ids []int) []TodoItem {
	var wg sync.WaitGroup
	todos := make([]TodoItem, len(ids))
	sem := make(chan struct{}, runtime.NumCPU())

	for i, id := range ids {
		id := id

		sem <- struct{}{} // acquire slot
		wg.Add(1)

		go func() {
			defer func() {
				<-sem // release slot
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
