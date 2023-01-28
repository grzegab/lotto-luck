package Lottery

import "grzegab.eu/lotto/Domain/Generator"

func CalculateCoverage[T Generator.LotteryNumbers | Generator.SystemNumbers, K Generator.LotteryNumbers | Generator.SystemNumbers](a T, b K) int {
	hits := 0

	for _, h1 := range a {
		for _, h2 := range b {
			if h1 == h2 {
				hits++
				break
			}
		}
	}

	return hits
}
