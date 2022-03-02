package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	//cannot be changed if it is a constant
	const conferenceTickets = 50
	var remainingTickets = 50

	//The printf function allows for placeholders and \n will create a new line when a string is written
	fmt.Printf("Welcome to our %v booking application <3\n", conferenceName)
	//this adds a space with the comma operend
	//fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "of these tickets remain available")
	fmt.Printf("We have a total of %v tickets and %v of these tickets remain available\n", conferenceTickets, remainingTickets)
	fmt.Println("Follow the below prompts in order to attend! We hope to see you there :)")

	//asks for username
	var userName string
	//the .= will auto assign its datatype, does not work with constants
	userTickets := 2

	userName = "Megan"

	fmt.Printf("User %v booked %v tickets for the %v \n", userName, userTickets, conferenceName)
	//this line displays the datatype the username variable
	//fmt.Printf("The username is type %T \n", userName)
}
