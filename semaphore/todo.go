package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

const todoBaseURL = "https://jsonplaceholder.typicode.com/todos"

type TodoItem struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func fetchTodo(ctx context.Context, todoID int) (TodoItem, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%d", todoBaseURL, todoID), nil)
	if err != nil {
		return TodoItem{}, fmt.Errorf("build request for todo %d: %w", todoID, err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return TodoItem{}, fmt.Errorf("execute request for todo %d: %w", todoID, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return TodoItem{}, fmt.Errorf("unexpected status %s for todo %d", res.Status, todoID)
	}

	var todo TodoItem
	if err := json.NewDecoder(res.Body).Decode(&todo); err != nil {
		return TodoItem{}, fmt.Errorf("decode todo %d: %w", todoID, err)
	}

	return todo, nil
}

func printTodos(todos []TodoItem) {
	if len(todos) == 0 {
		fmt.Println("No todos fetched.")
		return
	}

	sorted := append([]TodoItem(nil), todos...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].ID < sorted[j].ID
	})

	fmt.Printf("Fetched %d todo items\n", len(sorted))
	for _, todo := range sorted {
		fmt.Printf("  - #%d | user:%d | completed:%t | %s\n", todo.ID, todo.UserID, todo.Completed, todo.Title)
	}
}
