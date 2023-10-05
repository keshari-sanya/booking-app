package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for /* remainingTickets > 0 && len(bookings) <50 */ {
		//ask user for their name

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketsNumber {

			bookTickets(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName()
			fmt.Printf("The first names of the bookings are: %v\n\n", firstNames)

			if remainingTickets == 0 {
				//end the program
				fmt.Printf("Our conference is booked out. Come back next year!\n\n")
				break
			}
		} else if !isValidName || !isValidEmail || !isValidTicketsNumber {
			if !isValidName {
				fmt.Println("You have input wrong name")
			}
			if !isValidEmail {
				fmt.Println("You have input wrong email ID")
			}
			if !isValidTicketsNumber {
				fmt.Println("You have input incorrect ticket number")
			}

		} else {
			fmt.Printf("We only have %v tickets remaining, so you cannot book %v tickets\n\n", remainingTickets, userTickets)
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still remaining \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n\n")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email ID: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// use the struct function
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// create a map for a user
	/*
		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	*/

	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	fmt.Printf("\nList of bookings is %v\n\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confimation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n############")
	fmt.Printf("Sending ticket:\n %v \nto email address: %v\n", ticket, email)
	fmt.Printf("############\n")
	wg.Done()
}
