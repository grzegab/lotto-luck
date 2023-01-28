package DTO

type HitsMap map[int][7]int

type FunSystemStatsDto struct {
	count int
	hits  HitsMap
}

func NewFunSystemStatsDto() FunSystemStatsDto {
	f := FunSystemStatsDto{}

	return f
}

func (f *FunSystemStatsDto) SetCount(c int) {
	f.count = c
}

func (f *FunSystemStatsDto) GetCount() int {
	return f.count
}

func (f *FunSystemStatsDto) SetHits(h HitsMap) {
	f.hits = h
}

func (f *FunSystemStatsDto) GetHits() HitsMap {
	return f.hits
}
