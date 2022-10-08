package helpers

import (
	"booking_app_yt/src/models"
	"fmt"
)

func CreateGuestList(bookings *map[int]models.Booking) map[int]string {
	guestList := make(map[int]string)
	for key, value := range *bookings {
		guestList[key] = fmt.Sprintf("%s | f%d tickets |", value.Name, len(value.Tickets))
	}

	return guestList
}
