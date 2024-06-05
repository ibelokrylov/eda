package helpers

import (
	"math/rand"
	"time"
)

func GenerateCode() int {
	min := 1
	for i := 0; i < 5; i++ {
		min *= 10
	}
	max := min*10 - 1

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
