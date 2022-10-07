package services

import (
	hlp "booking_app_yt/src/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetUserInput(t string) string {
	var prompt = "\nHi, what's your name?"
	if t == "email" {
		prompt = "\nWhat's your email address?"
	}

	fmt.Printf("\n%s\n", prompt)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	userInput = strings.TrimSpace(userInput)

	if !hlp.ValidStrLen(userInput, 2) || t == "email" && !hlp.ValidEmail(userInput) {
		fmt.Println("\n\nInvalid input.\nNames must be longer than 2 letters, and emails must be correctly formatted.")
		return GetUserInput(t)
	}

	return userInput
}

func GetUserNumTickets() int {
	fmt.Println("\nHow many tickets would you like?")
	var userNumTickets int
	fmt.Scan(&userNumTickets)

	return userNumTickets
}
