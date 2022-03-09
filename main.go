package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

//to run go run .
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//the problem where the main thread ends before other threads complete can be fixed by
var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			//adds a thread so wg.add will run no matter when prog ends
			wg.Add(1)
			// the go makes it run concurrently in the background and will proceed to next task, will generate message in them iddle of next task
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("TuserTickets, firstName, lastName, emailhe first names of bookings are : %v\n", firstNames)
			if remainingTickets == 0 {
				//end program
				fmt.Printf("Our conference is booked out, come back next year!")
				break
			}

		} else if userTickets == remainingTickets {
			fmt.Printf("Success! %v %v booked %v tickets for the %v, Email sent to: %v \n", firstName, lastName, userTickets, conferenceName, email)
			fmt.Printf("You got the last tickets!")
			break

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			fmt.Println("youre input data is invalid")
		}
		//waits for goroutine to finish
		wg.Wait()
	}

	//switch statement practice
	city := "London"
	switch city {
	case "New York":
		//execute code for booking ny conference tickets
	case "Singapore":
		//execute code for booking singapore conference tickets
	case "London", "Paris":
		//execute code for booking singapore conference tickets
	case "Berlin":
		//execute code for booking singapore conference tickets
	case "Mexico City":
		//execute code for booking singapore conference tickets
	case "Hong Kong":
		//execute code for booking singapore conference tickets
	default:
		fmt.Println("No valid city selected")
	}

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application <3\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v of %v tickets remain available\n", conferenceName, remainingTickets, conferenceTickets)
	fmt.Println("Follow the below prompts in order to attend! We hope to see you there :)")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	//asks for firstName,lastName,email, and tickets
	var firstName string
	var lastName string
	var email string
	var userTickets uint
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)

	//create a map for a user, cannot mix data types
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Success! %v %v booked %v tickets for the %v, Email sent to: %v \n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("%v tickets are now remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("########\n")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Printf("########\n")
	wg.Done()
}
