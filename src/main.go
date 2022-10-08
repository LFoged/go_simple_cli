package main

import (
	"booking_app_yt/src/helpers"
	hlp "booking_app_yt/src/helpers"
	mdl "booking_app_yt/src/models"
	srv "booking_app_yt/src/services"
	"fmt"
	"time"
)

func main() {
	const conferenceName = "Shmazinga Bazinga"
	const maxTickets = 200

	bookings := make(map[int]mdl.Booking, 0)
	remainingTickets := maxTickets

	for remainingTickets > 0 {
		username := srv.GetUserInput("name")
		fmt.Printf("\n\nWelcome to the %s conference, %s!", conferenceName, username)
		fmt.Printf("\nWe have %d tickets left\n", remainingTickets)

		userTickets, userNumTickets := srv.IssueTickets(&remainingTickets)
		fmt.Println("\nTickets issued!")

		userEmail := srv.GetUserInput("email")
		bookingRef := helpers.GenRandInt("bookingRef")
		userBooking := mdl.Booking{
			BookingRef:   bookingRef,
			Name:         username,
			Email:        userEmail,
			Tickets:      userTickets,
			PurchaseDate: time.Now().String()}

		remainingTickets -= *userNumTickets
		bookings[bookingRef] = userBooking

		fmt.Printf("\n\nHere is your booking:\n%v.\n\nA confirmation has been sent to %s", userBooking, userEmail)
		fmt.Println("\n\n_______________________________________")
	}

	fmt.Println("\n\nWE'RE FULLY BOOKED!!!")

	fmt.Println("\n\n***************************")
	fmt.Printf("\n\nGUESTLIST: %v\n\n", hlp.CreateGuestList(&bookings))
	fmt.Println("***************************")

	fmt.Println("\n\n***************************")
	fmt.Printf("\n\nBOOKINGS: %v\n\n", bookings)
	fmt.Println("***************************")

}
