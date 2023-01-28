package MongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"grzegab.eu/lotto/Domain/Lotto"
	"time"
)

type ResultWithPrediction struct {
	Id               primitive.ObjectID `bson:"_id"`
	LottoID          int                `bson:"LottoID"`
	Date             time.Time          `bson:"Date"`
	Numbers          []int              `bson:"Numbers"`
	PredictionYearly []int              `bson:"PredictionYearly"`
	PredictionWhole  []int              `bson:"PredictionWhole"`
	OwnNumbers       []int              `bson:"OwnNumbers"`
	HitsYear         int                `bson:"HitsYear"`
	HitsWhole        int                `bson:"HitsWhole"`
	HitsOwn          int                `bson:"HitsOwn"`
}

func InsertLottoRecord(lwp Lotto.WithPrediction) error {
	c := OpenConnection(LottoTable)
	defer CloseConnection(c)

	bsonD := convertToBsonD(lwp)
	_, err := mongoCollection.InsertOne(ctx, bsonD)
	return err
}

func GetAllPredictionResults() []ResultWithPrediction {
	c := OpenConnection(LottoTable)
	defer CloseConnection(c)

	//get data
	filter := bson.D{{}}
	cursor, err := mongoCollection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var results []ResultWithPrediction
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	var response []ResultWithPrediction

	for _, result := range results {
		cursor.Decode(&result)
		response = append(response, result)
	}

	return response
}

func convertToBsonD(lwp Lotto.WithPrediction) bson.D {
	return bson.D{
		{"LottoID", lwp.LottoID},
		{"Date", lwp.Date},
		{"Numbers", lwp.Numbers},
		{"PredictionYearly", lwp.PredictionYearly},
		{"PredictionWhole", lwp.PredictionWhole},
		{"OwnNumbers", lwp.OwnNumbers},
		{"HitsYear", lwp.HitsYear},
		{"HitsWhole", lwp.HitsWhole},
		{"HitsOwn", lwp.HitsOwn},
	}
}
