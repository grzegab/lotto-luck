package Cli

import (
	"fmt"
	"grzegab.eu/lotto/Application/DTO"
)

func PrintResponse(dto DTO.FunResponseDto) {
	fmt.Println("")
	fmt.Println(fmt.Sprintf("There were %d hits until selected option", dto.HowManyUntil()))
	fmt.Println("")
	fmt.Println("Results saved to file:", dto.SavedFilename())
}

func PrintLottoResponse(dto DTO.LottoStats) {
	fmt.Println("")
	fmt.Println(fmt.Sprintf("There were %d hits ", dto.GetResultAffected()))
	fmt.Println("")
	fmt.Println("Results saved to file:", dto.GetSavedFilename())
}
