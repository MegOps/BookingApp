package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	//The printf function allows for placeholders and \n will create a new line when a string is written
	fmt.Printf("Welcome to our %v booking application <3\n", conferenceName)
	//this adds a space with the comma operend
	//fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "of these tickets remain available")
	fmt.Printf("We have a total of %v tickets and %v of these tickets remain available\n", conferenceTickets, remainingTickets)
	fmt.Println("Follow the below prompts in order to attend! We hope to see you there :)")

	for {

		//asks for firstName,lastName,email, and tickets
		var firstName string
		var lastName string
		var email string
		var userTickets int
		//fmt.Printf("User %v booked %v tickets for the %v \n", userName, userTickets, conferenceName)
		//this line displays the datatype the username variable
		//fmt.Printf("The username is type %T \n", userName)
		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email:")
		fmt.Scan(&email)
		fmt.Println("Please enter the amount of tickets you would like to purchase:")
		fmt.Scan(&userTickets)

		if userTickets >= int(remainingTickets) {
			fmt.Printf("We only have %v tickets remaining. So we are unable to book %v tickets\n", remainingTickets, userTickets)
			continue

		}
		fmt.Printf("Success! %v %v booked %v tickets for the %v, Email sent to: %v \n", firstName, lastName, userTickets, conferenceName, email)

		bookings = append(bookings, firstName+" "+lastName)

		remainingTickets = remainingTickets - uint(userTickets)
		fmt.Printf("%v tickets are now remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Printf("Our conference is booked out, come back next year!")
			break
		}

	}

}
