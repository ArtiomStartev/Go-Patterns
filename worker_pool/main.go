package main

import (
	"fmt"
)

func main() {
	// Create new tasks
	tasks := []Task{
		&EmailTask{Email: "email1@gmail.com", Subject: "Subject 1", MessageBody: "Message 1"},
		&ImageProcessingTask{ImageURL: "https://image1.com"},
		&EmailTask{Email: "email2@gmail.com", Subject: "Subject 2", MessageBody: "Message 2"},
		&ImageProcessingTask{ImageURL: "https://image2.com"},
		&EmailTask{Email: "email3@gmail.com", Subject: "Subject 3", MessageBody: "Message 3"},
		&ImageProcessingTask{ImageURL: "https://image3.com"},
		&EmailTask{Email: "email4@gmail.com", Subject: "Subject 4", MessageBody: "Message 4"},
		&ImageProcessingTask{ImageURL: "https://image4.com"},
		&EmailTask{Email: "email5@gmail.com", Subject: "Subject 5", MessageBody: "Message 5"},
		&ImageProcessingTask{ImageURL: "https://image5.com"},
	}

	// Create a worker pool with five concurrent workers.
	wp := NewWorkerPool(tasks, 5)

	// Run the pool
	wp.Run()
	fmt.Println("All tasks have been processed!")
}
