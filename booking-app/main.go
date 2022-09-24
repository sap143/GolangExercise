package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTicket int = 50

	fmt.Printf("conferenceTicket is %T, RemainingTicket is %T, ConferenceName %T,", conferenceTicket, remainingTicket, conferenceName)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available \n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")
	var bookings []string
	var firstName string
	var lastName string
	var email string
	var userTicket int
	// ask user for their FistName ,LastName,Email,Number of Ticket
	for {

		fmt.Println("Enter the FisrtName")
		fmt.Scan(&firstName)
		fmt.Println("Enter the LastName")
		fmt.Scan(&lastName)
		fmt.Println("Enter the email")
		fmt.Scan(&email)
		fmt.Println("Enter the how many ticket you want")
		fmt.Scan(&userTicket)
		bookings = append(bookings, firstName+" "+lastName)
		if userTicket > remainingTicket {
			fmt.Printf("We have only %v tickets remaining, so can`t book %v ticket\n", remainingTicket, userTicket)
			break
		}
		remainingTicket = remainingTicket - userTicket
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("The first names of booking are  :%v\n", firstNames)
		if remainingTicket == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year")
			break
		}
	}

}
