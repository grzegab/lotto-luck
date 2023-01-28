package FileSystem

import (
	"encoding/json"
	"os"
)

const FilePath = "var/startGames"

type Results struct {
	DrawDate     string `json:"drawDate"`
	DrawSystemId int    `json:"drawSystemId"`
	GameType     string `json:"gameType"`
	ResultsJson  []int  `json:"resultsJson"`
}

type Items struct {
	DrawDate             string    `json:"drawDate"`
	DrawSystemId         int       `json:"drawSystemId"`
	GameType             string    `json:"gameType"`
	IsNewEuroJackpotDraw bool      `json:"isNewEuroJackpotDraw"`
	MultiplierValue      int       `json:"multiplierValue"`
	Results              []Results `json:"results"`
}

type Games struct {
	TotalRows int     `json:"totalRows"`
	Code      int     `json:"code"`
	Items     []Items `json:"items"`
}

type LottoDraw struct {
	ID      int
	Date    string
	Numbers []int
}

func ReadLottoData() []LottoDraw {
	var lottoDraws []LottoDraw
	var bodyResult Games

	dat, _ := os.ReadFile(FilePath)
	json.Unmarshal(dat, &bodyResult)

	for _, i := range bodyResult.Items {
		if i.GameType == "Lotto" {
			for _, g := range i.Results {
				if g.GameType == "Lotto" {
					draw := LottoDraw{
						ID:      g.DrawSystemId,
						Date:    g.DrawDate,
						Numbers: g.ResultsJson,
					}
					lottoDraws = append([]LottoDraw{draw}, lottoDraws...)
				}
			}
		}
	}

	return lottoDraws
}
