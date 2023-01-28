package DTO

import "grzegab.eu/lotto/Domain/Lotto"

type LottoDto struct {
	LottoGames []Lotto.WithPrediction
}
