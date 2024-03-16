package randomnumbergenerator

import (
	"math/rand"
)

func RandomNumber() int {
	randomInt := rand.Intn(9000) + 1000
	return randomInt
}
