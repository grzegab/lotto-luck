package DrawSet

import (
	"math/rand"
	"time"
)

const numberPick = 6
const maxNumber = 49

type DrawNumbers [numberPick]int

func generateRandomNumberForDraw() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	number := r.Intn(maxNumber + 1)

	if number == 0 {
		return generateRandomNumberForDraw()
	}

	return number
}
