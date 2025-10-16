package main

import (
	"context"
	"fmt"
	"runtime"

	"golang.org/x/sync/errgroup"
)

func fetchTodosWithErrGroup(ctx context.Context, ids []int) []TodoItem {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(runtime.NumCPU())
	todos := make([]TodoItem, len(ids))

	for i, id := range ids {
		id := id
		g.Go(func() error {
			todo, err := fetchTodo(ctx, id)
			if err != nil {
				fmt.Println("Error fetching todo: ", err)
				return nil
			}

			todos[i] = todo
			return nil
		})
	}

	g.Wait()

	return todos
}
