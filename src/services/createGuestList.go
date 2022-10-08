package services

import (
	"booking_app_yt/src/models"
	"fmt"
)

func CreateGuestList(bookings *[]models.Booking) []string {
	guestList := []string{}
	for index, booking := range *bookings {
		guestList = append(guestList, fmt.Sprintf("%d-%s--%d_tickets_", index+1, booking.Name, len(booking.Tickets)))
	}

	return guestList
}
