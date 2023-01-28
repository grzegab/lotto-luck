package Application

import (
	"encoding/json"
	"fmt"
	"grzegab.eu/lotto/Application/DTO"
	"grzegab.eu/lotto/Domain/Generator"
	"grzegab.eu/lotto/Domain/Lottery"
	"grzegab.eu/lotto/Infrastructure/MongoDB"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func FunService(dto DTO.FunNumbersDto) DTO.FunResponseDto {
	response := DTO.FunResponseDto{}

	systemHits(dto.SystemHits())

	countUntil := howManyUntil(dto.GamesUntil())
	response.SetHowManyUntil(countUntil)

	filename := saveFunToFile()
	response.SetSavedFilename(filename)

	return response
}

func systemHits(untilHitsCount int) {
	var wg sync.WaitGroup

	draw := Lottery.NewDraw()
	system := Lottery.NewSystem()
	hits := system.CountSystemHits(draw)
	client := MongoDB.OpenConnection(MongoDB.FunSystemTable)
	ch := make(chan DTO.FunSystemDataDto)

	for w := 0; w <= untilHitsCount; w++ {
		for systemCount, hitNumber := range hits {
			wg.Add(1)
			go calcSystemHit(&wg, ch, systemCount, hitNumber, draw.Numbers, system.Numbers[:systemCount])
		}

		for i := 0; i <= len(hits)-1; i++ {
			wg.Add(1)
			go insertSystemHit(&wg, <-ch)
		}
	}

	wg.Wait()
	MongoDB.CloseConnection(client)
}

func howManyUntil(i int) int {
	client := MongoDB.OpenConnection(MongoDB.FunUntilHitTable)
	defer MongoDB.CloseConnection(client)

	d := Lottery.NewDraw()
	var winningNumbers Generator.LotteryNumbers
	count := 0
	for w := 0; w <= 1000000; w++ {
		count++
		l := Lottery.NewDraw()
		hits := Lottery.CalculateCoverage(d.Numbers, l.Numbers)

		if hits >= i {
			winningNumbers = l.Numbers
			break
		}
	}

	dto := DTO.NewFunUntilHitDto(i, count, d.Numbers, winningNumbers)
	insertUntilHit(dto)

	return count
}

func calcSystemHit(wg *sync.WaitGroup, ch chan DTO.FunSystemDataDto, systemCount, hitNumber int, d, s []int) {
	defer wg.Done()

	ch <- DTO.NewFunSystemDataDto(hitNumber, systemCount, d, s)
}

func insertSystemHit(wg *sync.WaitGroup, dto DTO.FunSystemDataDto) {
	defer wg.Done()

	if err := MongoDB.InsertFunSystemRecord(dto); err != nil {
		panic(err)
	}
}

func insertUntilHit(dto DTO.FunUntilHitDto) {
	if err := MongoDB.InsertFunUntilHitRecord(dto); err != nil {
		panic(err)
	}
}

func saveFunToFile() string {
	f, err := os.Create(filepath.Join(path, filepath.Base(funFilename)))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Fun System DATA
	funSystemData := MongoDB.GetFunSystemRecords()
	funSysCountText := fmt.Sprintf("const funSystemCount = %d\n", funSystemData.GetCount())
	if _, err := f.WriteString(funSysCountText); err != nil {
		log.Println(err)
	}

	jj, _ := json.Marshal(funSystemData.GetHits())
	funSysResults := fmt.Sprintf("const funSystemResults = %v\n", string(jj))
	if _, err := f.WriteString(funSysResults); err != nil {
		log.Println(err)
	}

	funUntilHitData := MongoDB.GetFunUntilHitRecords()
	funUntilHitCountText := fmt.Sprintf("const funUntilHitCount = %d\n", funUntilHitData.GetUntilHitGamesCount())
	if _, err := f.WriteString(funUntilHitCountText); err != nil {
		log.Println(err)
	}
	var jFunUntilCountTable = [6]int{}
	var jFunUntilMediumTable = [6]float64{}

	for l, v := range funUntilHitData.GetUntilHitGameForNumberCount() {
		cur := l - 1
		jFunUntilCountTable[cur] = v
	}

	for l, v := range funUntilHitData.GetMediumHits() {
		cur := l - 1
		jFunUntilMediumTable[cur] = v
	}

	jFunUntilCount, _ := json.Marshal(jFunUntilCountTable)
	jFunUntilMedium, _ := json.Marshal(jFunUntilMediumTable)
	funUntilHitResults := fmt.Sprintf(
		"const funUntilHitResults = {\"count\": %v, \"medium\": %v}\n",
		string(jFunUntilCount), string(jFunUntilMedium))
	if _, err := f.WriteString(funUntilHitResults); err != nil {
		log.Println(err)
	}

	return funFilename
}
