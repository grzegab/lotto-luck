package MongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"grzegab.eu/lotto/Application/DTO"
)

func InsertFunSystemRecord(dto DTO.FunSystemDataDto) error {
	_, err := mongoCollection.InsertOne(ctx, dto.BsonD())
	return err
}

func InsertFunUntilHitRecord(dto DTO.FunUntilHitDto) error {
	_, err := mongoCollection.InsertOne(ctx, dto.BsonD())
	return err
}

func GetFunSystemRecords() DTO.FunSystemStatsDto {
	c := OpenConnection(FunSystemTable)
	defer CloseConnection(c)

	fs := DTO.NewFunSystemStatsDto()
	count := 0

	//get data
	filter := bson.D{{}}
	cursor, err := mongoCollection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var results []DTO.FunSystemDataDto
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	system := DTO.HitsMap{
		7:  {0, 0, 0, 0, 0, 0, 0},
		8:  {0, 0, 0, 0, 0, 0, 0},
		9:  {0, 0, 0, 0, 0, 0, 0},
		10: {0, 0, 0, 0, 0, 0, 0},
		11: {0, 0, 0, 0, 0, 0, 0},
		12: {0, 0, 0, 0, 0, 0, 0},
	}

	for _, result := range results {
		cursor.Decode(&result)
		count++

		a := system[result.Size]
		a[result.Hits]++
		system[result.Size] = a
	}

	fs.SetCount(count)
	fs.SetHits(system)

	return fs
}

func GetFunUntilHitRecords() DTO.FunUntilHitStatsDto {
	c := OpenConnection(FunUntilHitTable)
	defer CloseConnection(c)

	fs := DTO.NewFunUntilHitStatsDto()
	filter := bson.D{{}}
	cursor, err := mongoCollection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	type UntilHitsNumbersMap map[int][]int

	var numbersMap = UntilHitsNumbersMap{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
	}

	untilHitMedium := DTO.UntilHitsMediumMap{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}

	untilHitCount := DTO.UntilHitsCountMap{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}

	var results []DTO.FunUntilHitDto
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	games := 0
	for _, result := range results {
		//Hits -> numer which was played
		//Count -> games until count number (Hits) was achieved
		cursor.Decode(&result)
		games++

		untilHitCount[result.Hits]++
		numbersMap[result.Hits] = append(numbersMap[result.Hits], result.Count)
	}

	//Calculate medium hits for number
	for g, k := range numbersMap {
		sum := 0
		for i := 0; i < len(k); i++ {
			sum += k[i]
		}

		med := 0.0
		if len(k) > 0 {
			med = float64(sum / len(k))
		}

		untilHitMedium[g] = med
	}

	fs.SetUntilHitGamesCount(games)
	fs.SetMediumHits(untilHitMedium)
	fs.SetUntilHitGameForNumberCount(untilHitCount)

	return fs
}
