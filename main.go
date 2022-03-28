package main

import (
	"fmt"
)

func introduction(name string, tot uint, rem uint) {
	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have a total of %v tickets and %v remaining\n", tot, rem)
	fmt.Println("Get your tickets here to attend")
}

func namePrompt() string {
	var guestFirstName string
	var guestLastName string
	fmt.Printf("First name: ")
	fmt.Scan(&guestFirstName)
	fmt.Printf("Last name: ")
	fmt.Scan(&guestLastName)
	return guestFirstName + " " + guestLastName
}

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	introduction(conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Guest's name is %s", namePrompt())
}
