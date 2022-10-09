// in golang projects everything get into packages
package main
// imports to provide extenrnal functionality
import "fmt"

// starting point of every application is main function
func main() {
	// varriables which were declared but not used are error

	// syntax sugar is using ':=' instead of 'var' 
	// (not working with consts and explicitly defined types)
	cinemaName := "cinema about Golang"
	// simple const
	const cinemaTickets = 50
	var remainingTickets uint = 50

	// to print types of variables type %T
	fmt.Printf("cinemaName type is %T\n", cinemaName)
	
	// fmt.Printf is for formating prints with variables
	fmt.Printf("Welcome to the %v booking application\n", cinemaName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", cinemaTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask users for their names by using pointer (address of var where it located)
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Amount of tickets to buy: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
}
