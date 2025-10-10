package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var tickets = 500
	ticketsChan := make(chan int)
	doneChan := make(chan struct{})

	go manageTickets(ticketsChan, doneChan, &tickets)

	for i := range [2000]int{} {
		wg.Add(1)
		go buyTicket(&wg, ticketsChan, i+1)
	}

	wg.Wait()
	doneChan <- struct{}{}

	time.Sleep(1 * time.Second)
	fmt.Println("Finished!")
}

func manageTickets(ticketsChan <-chan int, doneChan <-chan struct{}, remainingTickets *int) {
	for {
		select {
		case userId := <-ticketsChan:
			if *remainingTickets > 0 {
				*remainingTickets--
				fmt.Printf("User %d purchased a ticket. Tickets remaining: %d\n", userId, *remainingTickets)
			} else {
				fmt.Printf("User %d found no ticket.\n", userId)
			}
		case <-doneChan:
			fmt.Printf("Tickets remaining: %d\n", *remainingTickets)
			return
		}
	}
}

func buyTicket(wg *sync.WaitGroup, ticketsChan chan<- int, userId int) {
	defer wg.Done()
	ticketsChan <- userId
}
