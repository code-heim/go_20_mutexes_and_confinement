package main

import (
	"fmt"
	"sync"
)

/*
A function that is responsible for managing ticket allocations
and responding to user requests.
It listens for incoming requests on a channel (ticketChan) and signals
on another channel (doneChan) when it's time to stop.
*/
func manageTicket(ticketChan chan int, doneChan chan bool, tickets *int) {
	for {
		select {
		case user := <-ticketChan:
			if *tickets > 0 {
				*tickets--
				fmt.Printf("User %d purchased a ticket. Tickets remaining: %d\n",
					user, *tickets)
			} else {
				fmt.Printf("User %d found no tickets.\n", user)
			}
		case <-doneChan:
			fmt.Printf("Tickets remaining: %d\n", *tickets)
		}

	}
}

/*
A function that simulates a user trying to buy a ticket.
It sends a request to the manageTicket goroutine through ticketChan.
*/
func buyTicket(wg *sync.WaitGroup, ticketChan chan int, userId int) {
	defer wg.Done()
	ticketChan <- userId
}

func main() {
	var wg sync.WaitGroup        // WaitGroup to wait for all goroutines to finish
	tickets := 500               // Total number of tickets available
	ticketChan := make(chan int) // Channel for sending ticket purchase requests
	doneChan := make(chan bool)  // Channel for signaling the stop

	go manageTicket(ticketChan, doneChan, &tickets)

	for userId := 0; userId < 200; userId++ {
		wg.Add(1)
		go buyTicket(&wg, ticketChan, userId)
	}

	wg.Wait()
	doneChan <- true
}
