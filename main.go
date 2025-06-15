package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const totalTickets uint = 50

var Bookings = make([]userData, 0)
var conferenceName string = "Go Conference"
var remainingTickets uint = 50 //alternate way to declare variables

type userData struct {
	firstName string
	lastName  string
	emailID   string
	tickets   uint
}

func main() {

	greetFunction()
	for {
		// get userInput
		firstName, lastName, emailId, tickets := getUserInput()

		// validate user Input
		isValidName, isValidEmail, isValidTickets, err := ValidateUserInput(firstName, lastName, emailId, tickets, remainingTickets)

		if isValidEmail && isValidName && isValidTickets {
			// Book tickets
			bookTickets(firstName, lastName, emailId, tickets)
			go sendTicket(firstName, lastName, emailId, tickets)
			var firstNames = printFirstname()
			remainingTickets = remainingTickets - tickets
			fmt.Printf(" total Bookings are %v \nRemaining tickets are %v\n", firstNames, remainingTickets)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, Comeback next year")
				break
			}
		} else {
			fmt.Printf("ERROR : %v\n", err)
			if !isValidName {
				fmt.Println("Details : User name should be minimumm of 2 characters")
			}

			if !isValidEmail {
				fmt.Println("Details : Email ID should have @ in it")
			}

			if !isValidTickets {
				fmt.Printf("Details : Remaining tickets for booking are  %v\nPlease book within this\n\n\n\n", remainingTickets)
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
		firstNames = append(firstNames, boooking.firstName)
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
	var userData = userData{
		firstName: firstName,
		lastName:  lastName,
		emailID:   emailId,
		tickets:   tickets,
	}
	Bookings = append(Bookings, userData)
	// fmt.Printf("list of bookings are : %v", Bookings)
	fmt.Printf("Thank you %v %v for booking %v, you will receive confirmation at %v . \n", firstName, lastName, tickets, emailId)

}
func ValidateUserInput(firstName string, lastName string, emailId string, tickets uint, remainingTickets uint) (bool, bool, bool, error) {
	isValidName := (len(firstName) > 2) || (len(lastName) > 2)
	isValidEmail := strings.Contains(emailId, "@")
	isValidTickets := remainingTickets-tickets >= uint(0)
	if isValidName && isValidTickets && isValidEmail {
		return isValidName, isValidEmail, isValidTickets, nil
	} else {
		return isValidName, isValidEmail, isValidTickets, errors.New("Not a valid user input")
	}

}

func sendTicket(firstName string, lastName string, emailId string, tickets uint) {
	time.Sleep(time.Second * 10)
	var ticket = fmt.Sprintf("%v tickets for %v %v", tickets, firstName, lastName)

	fmt.Println("\n###########################################################")
	fmt.Printf("Sending tickets : \n %v \nto email ID : %v ", ticket, emailId)
	fmt.Println("###########################################################\n")

}
