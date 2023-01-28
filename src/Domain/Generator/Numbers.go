package Generator

import (
	"math/rand"
	"time"
)

const lotteryNumberCount = 6
const systemNumberCount = 12
const maxNumber = 49

type LotteryNumbers []int
type SystemNumbers []int

func GenerateLotteryNumbers() LotteryNumbers {
	var drawSet LotteryNumbers
	currentNumbers := 0

	for currentNumbers <= (lotteryNumberCount - 1) {
		number := GeneratePositiveRandomNumberUpTo(maxNumber)
		if !stackContainNumber(number, drawSet) {
			drawSet = append(drawSet, number)
			currentNumbers++
		}
	}

	return drawSet
}

func GenerateSystemNumbers() SystemNumbers {
	var drawSet SystemNumbers
	currentNumbers := 0

	for currentNumbers <= (systemNumberCount - 1) {
		number := GeneratePositiveRandomNumberUpTo(maxNumber)
		if !stackContainNumber(number, drawSet) {
			drawSet = append(drawSet, number)
			currentNumbers++
		}
	}

	return drawSet
}

func GeneratePositiveRandomNumberUpTo(n int) int {
	maxNumber := n + 1

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	number := r.Intn(maxNumber)

	if number == 0 {
		return GeneratePositiveRandomNumberUpTo(n)
	}

	return number
}

func stackContainNumber(n int, s []int) bool {
	for _, r := range s {
		if r == n {
			return true
		}
	}

	return false
}
