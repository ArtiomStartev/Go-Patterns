package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work executed by the worker pool.
type Task interface {
	Process()
}

// EmailTask sends an email with a subject and body.
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Process simulates the work involved in sending an email.
func (t *EmailTask) Process() {
	fmt.Printf("Sending email to %s...\n", t.Email)
	// Simulate a time-consuming task.
	time.Sleep(2 * time.Second)
}

// ImageProcessingTask fetches and processes an image.
type ImageProcessingTask struct {
	ImageURL string
}

// Process simulates the image processing work load.
func (t *ImageProcessingTask) Process() {
	fmt.Printf("Processing the image %s\n", t.ImageURL)
	// Simulate a time-consuming task.
	time.Sleep(5 * time.Second)
}

// WorkerPool orchestrates concurrent workers that execute pending tasks.
type WorkerPool struct {
	pendingTasks []Task
	tasksChan    chan Task
	workerCount  int
	wg           sync.WaitGroup
}

// NewWorkerPool builds a worker pool with the provided tasks and worker count.
func NewWorkerPool(tasks []Task, workerCount int) *WorkerPool {
	return &WorkerPool{
		pendingTasks: tasks,
		workerCount:  workerCount,
	}
}

// worker consumes tasks from the channel and signals completion.
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

// Run starts the workers and processes every pending task.
func (wp *WorkerPool) Run() {
	// Initialize the tasks channel with enough buffer to hold every task.
	wp.tasksChan = make(chan Task, len(wp.pendingTasks))

	// Start workers.
	for i := 0; i < wp.workerCount; i++ {
		go wp.worker()
	}

	// Send tasks to the channel.
	wp.wg.Add(len(wp.pendingTasks))
	for _, task := range wp.pendingTasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	// Wait for all tasks to finish.
	wp.wg.Wait()
}
