package Application

import (
	"encoding/json"
	"fmt"
	"os"
)

//tutaj sciagamy dane do bazy, tylko!!!

//czy DB jest czysta - jak tak laduj plik z danymini (na makas do roku 2022)
//sciagnij dane od 23 (querka 1000 wystarczy na to)
//informacja ze jak dlugo nie to reczny update i url z 10 000 i czyszczenie bazki
//(widac ze jest 6800 cos losowan, patrzec na ostatnie) niestety nie ma api

const URL = "https://www.lotto.pl/api/lotteries/draw-results/by-gametype?game=Lotto&index=%d&size=10&sort=drawSystemId&order=DESC"

const JSON = "{\"totalRows\":0,\"items\":[{\"drawSystemId\":6830,\"drawDate\":\"2023-01-12T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2023-01-12T22:00:00Z\",\"drawSystemId\":6830,\"gameType\":\"Lotto\",\"resultsJson\":[44,45,32,33,25,31],\"specialResults\":[]},{\"drawDate\":\"2023-01-12T22:00:00Z\",\"drawSystemId\":6830,\"gameType\":\"LottoPlus\",\"resultsJson\":[8,44,24,41,29,12],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false},{\"drawSystemId\":6829,\"drawDate\":\"2023-01-10T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2023-01-10T22:00:00Z\",\"drawSystemId\":6829,\"gameType\":\"Lotto\",\"resultsJson\":[40,3,48,10,30,11],\"specialResults\":[]},{\"drawDate\":\"2023-01-10T22:00:00Z\",\"drawSystemId\":6829,\"gameType\":\"LottoPlus\",\"resultsJson\":[1,9,18,6,21,20],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false},{\"drawSystemId\":6828,\"drawDate\":\"2023-01-07T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2023-01-07T22:00:00Z\",\"drawSystemId\":6828,\"gameType\":\"Lotto\",\"resultsJson\":[9,12,43,29,21,3],\"specialResults\":[]},{\"drawDate\":\"2023-01-07T22:00:00Z\",\"drawSystemId\":6828,\"gameType\":\"LottoPlus\",\"resultsJson\":[49,36,4,21,11,35],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false},{\"drawSystemId\":6827,\"drawDate\":\"2023-01-05T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2023-01-05T22:00:00Z\",\"drawSystemId\":6827,\"gameType\":\"Lotto\",\"resultsJson\":[31,43,49,45,7,21],\"specialResults\":[]},{\"drawDate\":\"2023-01-05T22:00:00Z\",\"drawSystemId\":6827,\"gameType\":\"LottoPlus\",\"resultsJson\":[3,6,21,17,1,23],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false},{\"drawSystemId\":6826,\"drawDate\":\"2023-01-03T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2023-01-03T22:00:00Z\",\"drawSystemId\":6826,\"gameType\":\"Lotto\",\"resultsJson\":[38,14,37,8,11,48],\"specialResults\":[]},{\"drawDate\":\"2023-01-03T22:00:00Z\",\"drawSystemId\":6826,\"gameType\":\"LottoPlus\",\"resultsJson\":[20,21,25,1,35,29],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false},{\"drawSystemId\":6825,\"drawDate\":\"2022-12-31T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2022-12-31T22:00:00Z\",\"drawSystemId\":6825,\"gameType\":\"Lotto\",\"resultsJson\":[36,46,47,16,29,19],\"specialResults\":[]},{\"drawDate\":\"2022-12-31T22:00:00Z\",\"drawSystemId\":6825,\"gameType\":\"LottoPlus\",\"resultsJson\":[6,13,12,44,22,45],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false},{\"drawSystemId\":6824,\"drawDate\":\"2022-12-29T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2022-12-29T22:00:00Z\",\"drawSystemId\":6824,\"gameType\":\"Lotto\",\"resultsJson\":[30,27,23,4,15,41],\"specialResults\":[]},{\"drawDate\":\"2022-12-29T22:00:00Z\",\"drawSystemId\":6824,\"gameType\":\"LottoPlus\",\"resultsJson\":[25,35,38,47,32,20],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false},{\"drawSystemId\":6823,\"drawDate\":\"2022-12-27T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2022-12-27T22:00:00Z\",\"drawSystemId\":6823,\"gameType\":\"Lotto\",\"resultsJson\":[9,11,28,2,41,17],\"specialResults\":[]},{\"drawDate\":\"2022-12-27T22:00:00Z\",\"drawSystemId\":6823,\"gameType\":\"LottoPlus\",\"resultsJson\":[45,31,42,43,11,4],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false},{\"drawSystemId\":6822,\"drawDate\":\"2022-12-24T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2022-12-24T22:00:00Z\",\"drawSystemId\":6822,\"gameType\":\"Lotto\",\"resultsJson\":[8,41,16,27,24,44],\"specialResults\":[]},{\"drawDate\":\"2022-12-24T22:00:00Z\",\"drawSystemId\":6822,\"gameType\":\"LottoPlus\",\"resultsJson\":[46,14,37,21,24,18],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false},{\"drawSystemId\":6821,\"drawDate\":\"2022-12-22T22:00:00Z\",\"gameType\":\"Lotto\",\"multiplierValue\":0,\"results\":[{\"drawDate\":\"2022-12-22T22:00:00Z\",\"drawSystemId\":6821,\"gameType\":\"Lotto\",\"resultsJson\":[5,3,38,35,2,46],\"specialResults\":[]},{\"drawDate\":\"2022-12-22T22:00:00Z\",\"drawSystemId\":6821,\"gameType\":\"LottoPlus\",\"resultsJson\":[21,30,47,34,48,6],\"specialResults\":[]}],\"showSpecialResults\":true,\"isNewEuroJackpotDraw\":false}],\"meta\":{},\"code\":200}"

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

func UpdateData() {

	getFileData()

	//getNewestData(1)
}

func getFileData() {
	dat, _ := os.ReadFile("var/startGames")

	var bodyResult Games

	json.Unmarshal(dat, &bodyResult)

	for _, i := range bodyResult.Items {
		if i.GameType == "Lotto" {

			for _, g := range i.Results {
				if g.GameType == "Lotto" {
					fmt.Println(g.ResultsJson)
				}
			}
		}
	}
}

func getNewestData(page int) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{
	//		CipherSuites: []uint16{
	//			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	//			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	//		},
	//		InsecureSkipVerify: true,
	//		MinVersion:         tls.VersionTLS12,
	//		MaxVersion:         tls.VersionTLS12,
	//	},
	//}
	//
	//client := &http.Client{Transport: tr}
	//req, err := http.NewRequest("GET", URL, nil)
	//if err != nil {
	//	panic(err)
	//}
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	//req.Header.Set("Access-Control-Allow-Origin", "*")
	//req.Header.Add("Accept-Language", "pl-PL,pl;q=0.9,en-US;q=0.8,en;q=0.7")
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//body, err := io.ReadAll(resp.Body)

	j := []byte(JSON)

	//fmt.Println(body)
	var bodyResult Games

	json.Unmarshal(j, &bodyResult)

	for _, i := range bodyResult.Items {
		if i.GameType == "Lotto" {
			for _, g := range i.Results {
				if g.GameType == "Lotto" {
					fmt.Println(g.ResultsJson)
				}
			}
		}
	}

	//if err != nil {
	//	panic(err)
	//}

	//bodyString := string(bodyResult)

}

func parseResponseObjects() {

}

func getCurrentDateFromDB() {

}

func checkIfInDB() {

	//sprawdz date z bazki
	//jesli pobrany obiekt data > bazka-data zpisz do db
	//jelsi data juz jest w bazce - przerwij wszystko, jest aktualne
	//zapisz ile wpisow dodanych
}

func ReadGamesFile() {

}
