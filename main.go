package main

import (
	"fmt"
	"strconv"
	"strings"
)

const totalTickets uint = 50

var Bookings = make([]map[string]string, 1)
var conferenceName string = "Go Conference"
var remainingTickets uint = 50 //alternate way to declare variables

func main() {

	greetFunction()
	for {
		// get userInput
		firstName, lastName, emailId, tickets := getUserInput()

		// validate user Input
		isValidName, isValidEmail, isValidTickets := ValidateUserInput(firstName, lastName, emailId, tickets, remainingTickets)
		if isValidEmail && isValidName && isValidTickets {
			// Book tickets
			bookTickets(firstName, lastName, emailId, tickets)
			var firstNames = printFirstname()
			remainingTickets = remainingTickets - tickets
			fmt.Printf(" total Bookings are %v \nRemaining tickets are %v\n", firstNames, remainingTickets)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, Comeback next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("ERROR : User name should be minimumm of 2 characters")
			}

			if !isValidEmail {
				fmt.Println("ERROR : Email ID should have @ in it")
			}

			if !isValidTickets {
				fmt.Printf("ERROR : Remaining tickets for booking are  %v\nPlease book within this\n\n\n\n", remainingTickets)
			}
		}

	}

}

func greetFunction() {
	fmt.Println("Welcome to", conferenceName)
	fmt.Printf("Total tickets are %v and remaining tickets are %v\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
}

func printFirstname() []string {
	firstNames := []string{}
	for _, boooking := range Bookings {
		firstNames = append(firstNames, boooking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailId string
	var tickets uint
	println("Enter your first name")
	fmt.Scan(&firstName)
	println("Enter your last name")
	fmt.Scan(&lastName)
	println("Enter your emailID")
	fmt.Scan(&emailId)
	println("Enter your tickets")
	fmt.Scan(&tickets)

	return firstName, lastName, emailId, tickets
}

func bookTickets(firstName string, lastName string, emailId string, tickets uint) {
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["emailID"] = emailId
	userData["noOfTickets"] = strconv.FormatUint(uint64(tickets), 10)
	Bookings = append(Bookings, userData)
	fmt.Printf("list of bookings are : %v", Bookings)
	// fmt.Printf("Thank you %v %v for booking %v, you will receive confirmation at %v . \n", firstName, lastName, tickets, emailId)

}
func ValidateUserInput(firstName string, lastName string, emailId string, tickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := (len(firstName) > 2) || (len(lastName) > 2)
	isValidEmail := strings.Contains(emailId, "@")
	isValidTickets := remainingTickets-tickets >= uint(0)

	return isValidName, isValidEmail, isValidTickets
}
