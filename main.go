package main

import (
	"booking_app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceName = "Go Conference"
const conferenceTickets = 50

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

	greetUsers()

	for {

		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTicket, firstName, lastName, email)

			wg.Add(1)
			go sendTickets(userTicket, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of booking are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")

			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign ")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number you have entered is incorrect")
			}
			fmt.Printf("Your imput data is invalid,try again\n")

		}
	}
	wg.Wait()
}
func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n ", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining Tickets are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
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
	var userTicket uint
	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)
	return firstName, lastName, email, userTicket
}
func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//create a map fo user

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf(" %v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(2 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending tickets:\n %v to email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}
