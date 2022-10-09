package helper

import "strings"

// to make func scope global we should use capitalized name
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	
	// you can return any number of values
	return isValidName, isValidEmail, isValidTicketNumber
}