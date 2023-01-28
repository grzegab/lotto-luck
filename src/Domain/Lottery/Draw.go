package Lottery

import (
	"grzegab.eu/lotto/Domain/Generator"
	"time"
)

type Draw struct {
	Numbers  Generator.LotteryNumbers
	GameDate time.Time
}

func NewDraw() Draw {
	d := Draw{}
	d.GameDate = time.Now()
	d.Numbers = Generator.GenerateLotteryNumbers()

	return d
}

func (d Draw) hits(n Draw) int {

	return 6
}

func (d Draw) containsNumber(n int) bool {
	for _, r := range d.Numbers {
		if r == n {
			return true
		}
	}

	return false
}

func (d Draw) getNumbers() Generator.LotteryNumbers {
	return d.Numbers
}
