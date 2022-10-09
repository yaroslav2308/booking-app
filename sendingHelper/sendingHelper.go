package sendingHelper

import (
	"fmt"
	"time"
)

// time-consuming process to show how concurrency work in Golang (simulation)
func SendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, email)
	fmt.Println("####################")
}
