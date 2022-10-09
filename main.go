// in golang projects everything get into packages
package main

// imports to provide extenrnal functionality
import (
	"booking-app/helper"
	"booking-app/sendingHelper"
	"fmt"
	// "sync"
	// "strconv"
	// "strings"
)

// to print types of variables type %T
// fmt.Printf("cinemaName type is %T\n", cinemaName)

// varriables which were declared but not used are error

// syntax sugar is using ':=' instead of 'var'
// (not working with consts and explicitly defined types)
const cinemaTickets = 50
var cinemaName = "cinema about Golang"
var remainingTickets uint = 50

// arrays (always fixed size)
// var bookings [50]string

// slice is below
var bookings = make([]UserData, 0)

// struct
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

// waitGroup 
// var wg = sync.WaitGroup {}

// starting point of every application is main function
func main() {
	
	greetUsers()

	// infinite for loop to book tickets 
	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		
		if !isValidName || !isValidEmail || !isValidTicketNumber  {
			if !isValidName {
				fmt.Println("Your name / last name lenght must be at least 2 characters")
			}

			if !isValidEmail {
				fmt.Println("Your email address must contain '@' symbol")
			}

			if !isValidTicketNumber {
				fmt.Println("Your ticket number must be less or equal to remaining tickets number")
			}
			fmt.Println("Try one more time")
			continue
		}

		bookTicket(userTickets, firstName, lastName, email)

		// create thread by using 'go' keyword
		go sendingHelper.SendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("All bokings: %v\n", firstNames)

		// if statement
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our cinema is booked out. Come back next week")
			break
		}
	}
}

func greetUsers() {
	// fmt.Printf is for formating prints with variables
	fmt.Printf("Welcome to the %v booking application\n", cinemaName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", cinemaTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	// for loop always need 2 values (index and element)
	// we do not use index so it turns into underscore '_'
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

	// ask users for their names by using pointer (address of var where it located)
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Amount of tickets to buy: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	
	// create a map for user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, cinemaName)
}


