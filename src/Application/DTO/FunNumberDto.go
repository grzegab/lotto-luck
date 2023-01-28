package DTO

type FunNumbersDto struct {
	gamesUntil int
	systemHits int
}

func (f *FunNumbersDto) GamesUntil() int {
	return f.gamesUntil
}

func (f *FunNumbersDto) SetGamesUntil(gamesUntil int) {
	f.gamesUntil = gamesUntil
}

func (f *FunNumbersDto) SystemHits() int {
	return f.systemHits
}

func (f *FunNumbersDto) SetSystemHits(systemHits int) {
	f.systemHits = systemHits
}
