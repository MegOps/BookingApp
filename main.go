package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	//cannot be changed if it is a constant
	const conferenceTickets = 50
	var remainingTickets uint = 50

	//The printf function allows for placeholders and \n will create a new line when a string is written
	fmt.Printf("Welcome to our %v booking application <3\n", conferenceName)
	//this adds a space with the comma operend
	//fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "of these tickets remain available")
	fmt.Printf("We have a total of %v tickets and %v of these tickets remain available\n", conferenceTickets, remainingTickets)
	fmt.Println("Follow the below prompts in order to attend! We hope to see you there :)")

	//asks for firstName,lastName,email, and tickets
	var firstName string
	var lastName string
	var email string
	var userTickets int

	//the .= will auto assign its datatype, does not work with constants
	//userName = "Megan"

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
	fmt.Printf("Success! %v %v booked %v tickets for the %v, Email sent to: %v \n", firstName, lastName, userTickets, conferenceName, email)

	remainingTickets = remainingTickets - uint(userTickets)
	fmt.Printf("%v tickets are now remaining for %v\n", remainingTickets, conferenceName)
	//when i add a number its an array
	var bookings []string
	//bookings[0] = firstName + " " + lastName
	//this makes a slice
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The array type : %T\n", bookings)
	fmt.Printf("The array length: %v\n", len(bookings))
}

//1:11:12 golang video
