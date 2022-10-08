package services

import (
	hlp "booking_app_yt/src/helpers"
	mdl "booking_app_yt/src/models"
	"fmt"
	"sync"
	"time"
)

func PrintBookings(bookings *map[int]mdl.Booking, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("\nPrinting guestlist...")
	time.Sleep(2 * time.Second)
	fmt.Println("\n\n***************************")
	fmt.Printf("\n\nGUESTLIST: %v\n\n", hlp.CreateGuestList(bookings))
	fmt.Println("***************************")

	time.Sleep(1 * time.Second)
	fmt.Println("\nPrinting bookings...")
	for k, v := range *bookings {
		wg.Add(1)
		time.Sleep(2 * time.Second)
		fmt.Println("\n\n***************************")
		fmt.Printf("\nBOOKING: %d\nName:%s | Email: %s\nTickets:%v\n", k, v.Name, v.Email, v.Tickets)
		fmt.Println("\n***************************")
		wg.Wait()
	}
}
