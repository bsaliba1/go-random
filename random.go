package random

import (
	"math/rand"
)

func RandomNumber() int {
    return rand.Intn(100)
}
