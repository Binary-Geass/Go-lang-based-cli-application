package main

import "fmt"

func main() {
	conferenceName := "Go Conferene"
	const conferenceTickets int = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are remaining\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend")
}
