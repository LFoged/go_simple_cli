package services

import (
	mdl "booking_app_yt/src/models"
	"fmt"
	"sync"
	"time"
)

func SendEmail(bookings *map[int]mdl.Booking, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("\n\nSending confirmation emails...")

	time.Sleep(20 * time.Second)

	for _, item := range *bookings {
		wg.Add(1)
		go time.Sleep(2 * time.Second)
		fmt.Printf("\n\nTo %s:\nHi, %s, here are your tickets!\nEnjoy!\n %v\n", item.Email, item.Name, item.Tickets)
		wg.Wait()
	}

}
