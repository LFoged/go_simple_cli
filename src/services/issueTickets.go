package services

import (
	hlp "booking_app_yt/src/helpers"
	"fmt"
)

func IssueTickets(remainingTickets *int) ([]int, *int) {
	userNumTickets := GetUserNumTickets()

	if userNumTickets < 1 || userNumTickets > *remainingTickets {
		fmt.Printf("\nWe have %d tickets left.\nChoose a number between 1-%d.", *remainingTickets, *remainingTickets)
		return IssueTickets(remainingTickets)
	}

	userTickets := hlp.CreateTickets(&userNumTickets)

	return userTickets, &userNumTickets
}
