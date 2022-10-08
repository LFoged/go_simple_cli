package helpers

func CreateTickets(numTickets *int) []int {
	userTickets := []int{}
	for i := 0; i < *numTickets; i++ {
		userTickets = append(userTickets, GenRandInt("ticket"))
	}

	return userTickets
}
