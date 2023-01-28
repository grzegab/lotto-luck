package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:root@database:27017/?maxPoolSize=20&w=majority"
const LottoTable = "lotto"
const FunSystemTable = "funSystem"
const FunUntilHitTable = "funUntil"

var mongoCollection *mongo.Collection
var database *mongo.Database
var client *mongo.Client
var ctx = context.TODO()

func OpenConnection(collection string) *mongo.Client {
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	client = c
	mongoCollection = client.Database("lotto").Collection(collection)

	return c
}

func CloseConnection(c *mongo.Client) {
	if err := c.Disconnect(ctx); err != nil {
		panic(err)
	}
}
