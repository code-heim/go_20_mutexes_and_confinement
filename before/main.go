package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func buyTicket(wg *sync.WaitGroup, userId int, remainingTickets *int) {
	defer wg.Done()
	if *remainingTickets > 0 {
		*remainingTickets-- // User purchases a ticket
		fmt.Printf("User %d purchased a ticket. Tickets remaining: %d\n",
			userId, *remainingTickets)

	} else {
		fmt.Printf("User %d found no ticket.\n", userId)
	}
}

func main() {
	var tickets int = 500

	var wg sync.WaitGroup

	// Simulating a lot of users trying to buy tickets
	// For simplicity, let's assume number of tickets a user can buy is one
	for userId := 0; userId < 2000; userId++ { // More requests than tickets
		wg.Add(1)
		// Buy ticket for the user with ID userId
		go buyTicket(&wg, userId, &tickets)
	}

	wg.Wait()

}

