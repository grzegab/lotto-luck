package Cli

import (
	"fmt"
	"grzegab.eu/lotto/Application"
	"grzegab.eu/lotto/Application/DTO"
	"os"
)

func SelectGame() {
	printIntro()
	printMenu()
	menuSelection()
}

func printIntro() {
	fmt.Println("/=---------------------------------------=\\")
	fmt.Println("/ Small example of random lotto math game \\")
	fmt.Println("/=---------------------------------------=\\")
}

func printMenu() {
	fmt.Println("")
	fmt.Println("Select desire option:")
	fmt.Println("1. Calculate results for standard 6 digit")
	fmt.Println("2. Fun generator, check for possible things")
	fmt.Println("0. Exit")
	fmt.Println("--------------------------------------------")
}

func moreInfo() (int, int) {
	luckyGuessNumber := 5
	systemHits := 100000

	fmt.Println("")
	fmt.Println("Select more options:")

	fmt.Print("How many games until count happen (1-6, recommended 5): ")
	fmt.Scanf("%v", &luckyGuessNumber)

	fmt.Print("How many system hits (recommended 100000): ")
	fmt.Scanf("%v", &systemHits)

	return luckyGuessNumber, systemHits
}

func menuSelection() {
	var selectedNumber int
	fmt.Print("Your selection: ")
	fmt.Scanf("%v", &selectedNumber)

	switch selectedNumber {
	case 1:
		stats := Application.RunLottoStats()

		PrintLottoResponse(stats)
	case 2:
		luckyGuessNumber, systemHits := moreInfo()

		dto := DTO.FunNumbersDto{}
		dto.SetGamesUntil(luckyGuessNumber)
		dto.SetSystemHits(systemHits)
		response := Application.FunService(dto)

		PrintResponse(response)
	case 0:
		fmt.Println("EXIT!")
	default:
		fmt.Println("Wrong selection...")
		os.Exit(1)
	}

	os.Exit(0)
}
