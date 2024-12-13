package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string = []string{}

const conferenceTickets = 50

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		isValidInput := isValidName && isValidEmail && isValidTicketNumber

		if isValidInput {
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			firstNames := getFirstName(bookings)
			fmt.Printf("Tickets booked by %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come next year.")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("Enter your name correctly")
			}
			if !isValidEmail {
				fmt.Println("Enter a valid email address")
			}
			if !isValidTicketNumber {
				fmt.Println("Enter a valid number of tickets")
			}
			fmt.Printf("Your input data is invalid. Try again.")
		}

	}
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Printf("Enter your first name: ")
	_, _ = fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	_, _ = fmt.Scan(&lastName)
	fmt.Printf("Enter your email address: ")
	_, _ = fmt.Scan(&email)
	fmt.Printf("Enter number of tickets: ")
	_, _ = fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets, userTickets uint, bookings []string, firstName, lastName, email, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
