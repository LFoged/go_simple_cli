package services

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func createTickets(numTickets *int) []string {
	userTickets := []string{}
	for i := 0; i < *numTickets; i++ {
		rand.Seed(time.Now().UnixNano())
		ticketId := rand.Intn(9999999)
		userTickets = append(userTickets, strconv.Itoa(ticketId))
	}

	return userTickets
}

func IssueTickets(remainingTickets *int) ([]string, *int) {
	userNumTickets := GetUserNumTickets()

	if userNumTickets < 1 || userNumTickets > *remainingTickets {
		fmt.Printf("\nWe have %d tickets left.\nChoose a number between 1-%d.", *remainingTickets, *remainingTickets)
		return IssueTickets(remainingTickets)
	}

	userTickets := createTickets(&userNumTickets)

	return userTickets, &userNumTickets
}
