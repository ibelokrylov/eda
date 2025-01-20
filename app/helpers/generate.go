package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandomInt64() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int64(r.Uint64())
}
