package models

type Booking struct {
	BookingRef   int
	Name         string
	Email        string
	Tickets      []int
	PurchaseDate string
}
