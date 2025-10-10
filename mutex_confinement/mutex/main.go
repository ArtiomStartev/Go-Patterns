package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var tickets = 500

	for i := range [2000]int{} {
		wg.Add(1)
		go buyTicket(&wg, &mu, i+1, &tickets)
	}
	wg.Wait()
}

func buyTicket(wg *sync.WaitGroup, mu *sync.Mutex, userId int, remainingTickets *int) {
	defer wg.Done()
	mu.Lock()
	if *remainingTickets > 0 {
		*remainingTickets--
		fmt.Printf("User %d purchased a ticket. Tickets remaining: %d\n", userId, *remainingTickets)
	} else {
		fmt.Printf("User %d found no ticket.\n", userId)
	}
	mu.Unlock()
}
