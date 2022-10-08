package services

import (
	hlp "booking_app_yt/src/helpers"
	mdl "booking_app_yt/src/models"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func PrintBookings(bookings *map[int]mdl.Booking) {
	fmt.Println("\nPrinting guestlist...")
	time.Sleep(2 * time.Second)
	fmt.Println("\n\n***************************")
	fmt.Printf("\n\nGUESTLIST: %v\n\n", hlp.CreateGuestList(bookings))
	fmt.Println("***************************")

	time.Sleep(1 * time.Second)
	fmt.Println("\nPrinting bookings...")
	for k, v := range *bookings {
		time.Sleep(2 * time.Second)
		fmt.Println("\n\n***************************")
		fmt.Printf("\nBOOKING: %d\nName:%s | Email: %s\nTickets:%v\n", k, v.Name, v.Email, v.Tickets)
		fmt.Println("\n***************************")

	}

	wg.Done()
}
