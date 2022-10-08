package services

import (
	mdl "booking_app_yt/src/models"
	"fmt"
	"time"
)

func SendEmail(bookings *map[int]mdl.Booking) {
	fmt.Println("\n\nSending confirmation emails...")

	time.Sleep(20 * time.Second)

	for _, item := range *bookings {
		time.Sleep(1 * time.Second)
		fmt.Printf("\n\nTo %s:\nHi, %s, here are your tickets!\nEnjoy!\n %v\n", item.Email, item.Name, item.Tickets)
	}
	wg.Done()

}
