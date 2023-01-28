package Application

import (
	"fmt"
	"grzegab.eu/lotto/Application/DTO"
	"grzegab.eu/lotto/Domain/Lotto"
	"grzegab.eu/lotto/Infrastructure/FileSystem"
	"grzegab.eu/lotto/Infrastructure/MongoDB"
	"log"
	"os"
	"path/filepath"
	"time"
)

func RunLottoStats() DTO.LottoStats {
	stats := DTO.LottoStats{}

	resultsAffected := updateLottoData()
	savedFilename := saveLottoToFile()

	stats.SetResultAffected(resultsAffected)
	stats.SetSavedFilename(savedFilename)

	return stats
}

func updateLottoData() int {
	lottoData := FileSystem.ReadLottoData()
	resultCount := 0
	luckyNumbers := []int{5, 7, 13, 21, 36, 41}

	for _, g := range lottoData {
		var lottoWithPrediction = Lotto.WithPrediction{
			LottoID: g.ID,
			Date:    g.Date,
			Numbers: g.Numbers,
		}

		//make prediction for game based on last lottoData
		year, whole := getPrediction(lottoWithPrediction)
		lottoWithPrediction.PredictionYearly = year
		lottoWithPrediction.PredictionWhole = whole

		//set own numbers
		lottoWithPrediction.OwnNumbers = luckyNumbers

		//calculate hits for prediction
		lottoWithPrediction.CalculateHits()

		//insert into DB
		err := MongoDB.InsertLottoRecord(lottoWithPrediction)
		if err != nil {
			panic(err)
		}

		resultCount++
	}

	return resultCount
}

func getPrediction(l Lotto.WithPrediction) ([]int, []int) {
	var yearResults []int
	var wholeResults []int
	var yearNumberCount [50]int
	var wholeNumberCount [50]int

	results := MongoDB.GetAllPredictionResults()

	//return default values if first entry
	if len(results) == 1 {
		return []int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}
	}

	for u := 0; u < len(results)-1; u++ {
		checkDate, err := time.Parse("2006-01-02T15:04:05Z", l.Date)
		yearBefore := checkDate.AddDate(-1, 0, 0)
		//yearAfter := checkDate.AddDate(+1, 0, 0)
		lotteryDate := results[u].Date

		if err != nil {
			panic(err)
		}

		for _, n := range results[u].Numbers {
			wholeNumberCount[n]++
			if yearBefore.After(lotteryDate) {
				yearNumberCount[n]++
			}
		}
	}

	for lottoLeng := 0; lottoLeng <= 5; lottoLeng++ {
		maxNum := 0
		maxIndex := 0
		for y, o := range yearNumberCount {
			if o >= maxNum {
				maxNum = o
				maxIndex = y
			}
		}
		yearNumberCount[maxIndex] = 0
		yearResults = append(yearResults, maxIndex)
	}

	for lottoLeng := 0; lottoLeng <= 5; lottoLeng++ {
		maxNum := 0
		maxIndex := 0
		for y, o := range wholeNumberCount {
			if o > maxNum {
				maxNum = o
				maxIndex = y
			}
		}
		wholeNumberCount[maxIndex] = 0
		wholeResults = append(wholeResults, maxIndex)
	}

	return yearResults, wholeResults
}

func saveLottoToFile() string {
	fmt.Println("Saving to file", lottoFilename)
	f, err := os.Create(filepath.Join(path, filepath.Base(lottoFilename)))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Fun System DATA
	lottoData := MongoDB.GetAllPredictionResults()

	if _, err := f.WriteString("const dataSet = [\n"); err != nil {
		log.Println(err)
	}

	for _, v := range lottoData {
		string := fmt.Sprintf("['%v','%v','%v','%v','%v', '%v', '%v', '%v', '%v'],\n", v.LottoID, v.Date, v.Numbers, v.PredictionWhole, v.HitsWhole, v.PredictionYearly, v.HitsYear, v.OwnNumbers, v.HitsOwn)
		if _, err := f.WriteString(string); err != nil {
			log.Println(err)
		}
	}

	if _, err := f.WriteString("];\n"); err != nil {
		log.Println(err)
	}

	return lottoFilename
}
