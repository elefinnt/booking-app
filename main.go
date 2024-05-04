package main

import "fmt"

func main() {
	 conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking system!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for their name
fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets



	userTickets = 2
	fmt.Printf("Thank you %v %v for booking %v tickets to %v. Your tickets will be sent to %v\n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}