package helpers

import (
	"math/rand"
	"time"
)

func GenRandInt(t string) int {
	maxNum := 999999
	if t == "bookingRef" {
		maxNum = (9999)
	}
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(maxNum)
}
