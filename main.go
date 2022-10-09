// in golang projects everything get into packages
package main
// imports to provide extenrnal functionality
import (
	"fmt"
	"strings"
)

// starting point of every application is main function
func main() {
	// varriables which were declared but not used are error

	// syntax sugar is using ':=' instead of 'var' 
	// (not working with consts and explicitly defined types)
	cinemaName := "cinema about Golang"
	// simple const
	const cinemaTickets = 50
	var remainingTickets uint = 50
	// arrays (always fixed size)
	// var bookings [50]string
	
	// slice is below
	var bookings []string

	// to print types of variables type %T
	// fmt.Printf("cinemaName type is %T\n", cinemaName)
	
	// fmt.Printf is for formating prints with variables
	fmt.Printf("Welcome to the %v booking application\n", cinemaName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", cinemaTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// infinite for loop to book tickets 
	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		
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
			continue
		}
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, cinemaName)
	
		firstNames := []string{}

		// for loop always need 2 values (index and element)
		// we do not use index so it turns into underscore '_'
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("All bokings: %v\n", firstNames)

		// if statement
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our cinema is booked out. Come back next week")
			break
		}
	}

	
}
