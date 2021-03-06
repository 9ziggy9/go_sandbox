package main

import (
	"fmt"
)

func introduction(name string, tot uint, rem uint) {
	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have a total of %v tickets and %v remaining\n", tot, rem)
	fmt.Println("Get your tickets here to attend")
}

func namePrompt() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var tickets uint
	fmt.Printf("First name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Email: ")
	fmt.Scan(&email)
	fmt.Printf("Tickets: ")
	fmt.Scan(&tickets)
	return firstName, lastName, email, tickets
}

func bookUser(tot uint, rem *uint, bookings []string) {
	var userFName, userLName, userEmail, userTickets = namePrompt()
	fmt.Printf("%s %s with email %s ordered %d tickets. ", userFName,
		userLName, userEmail, userTickets)
	*rem = tot - userTickets
	for i := uint(0); i < userTickets; i++ {
		bookings[i] = userLName + ", " + userFName
	}
}

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	bookings := make([]string, 50)

	introduction(conferenceName, conferenceTickets, remainingTickets)
	bookUser(conferenceTickets, &remainingTickets, bookings)
	fmt.Println(remainingTickets, "are now remaining of", conferenceTickets)

	for _, c := range bookings {
		if c != "" {
			fmt.Println(c)
		}
	}
}
