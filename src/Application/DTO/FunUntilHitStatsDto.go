package DTO

type UntilHitsMediumMap map[int]float64
type UntilHitsCountMap map[int]int

type FunUntilHitStatsDto struct {
	count  int
	games  UntilHitsCountMap
	medium UntilHitsMediumMap
}

func NewFunUntilHitStatsDto() FunUntilHitStatsDto {
	f := FunUntilHitStatsDto{}

	return f
}

func (f *FunUntilHitStatsDto) SetUntilHitGamesCount(c int) {
	f.count = c
}

func (f *FunUntilHitStatsDto) GetUntilHitGamesCount() int {
	return f.count
}

func (f *FunUntilHitStatsDto) SetUntilHitGameForNumberCount(g UntilHitsCountMap) {
	f.games = g
}

func (f *FunUntilHitStatsDto) GetUntilHitGameForNumberCount() UntilHitsCountMap {
	return f.games
}

func (f *FunUntilHitStatsDto) SetMediumHits(h UntilHitsMediumMap) {
	f.medium = h
}

func (f *FunUntilHitStatsDto) GetMediumHits() UntilHitsMediumMap {
	return f.medium
}
