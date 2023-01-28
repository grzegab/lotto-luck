package Lottery

import (
	"grzegab.eu/lotto/Domain/Generator"
	"time"
)

type SystemHits map[int]int
type SystemNumbers map[int]int

type System struct {
	Numbers        []int
	GenerationDate time.Time
}

func NewSystem() System {
	s := System{
		Numbers:        Generator.GenerateSystemNumbers(),
		GenerationDate: time.Now(),
	}

	return s
}

func (s System) CountSystemHits(d Draw) SystemHits {
	sh := SystemHits{}
	hits := 0
	for i, h1 := range s.Numbers {
		for _, h2 := range d.Numbers {
			if h1 == h2 {
				hits++
				break
			}
		}
		if i >= 6 {
			sh[i+1] = hits
		}
	}

	return sh
}
