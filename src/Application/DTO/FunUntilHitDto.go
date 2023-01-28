package DTO

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FunUntilHitDto struct {
	Id          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	Hits        int                `bson:"hits"`
	Count       int                `bson:"count"`
	OriginalSet []int              `bson:"original"`
	UserSet     []int              `bson:"custom"`
}

func NewFunUntilHitDto(h int, c int, o []int, u []int) FunUntilHitDto {
	f := FunUntilHitDto{Hits: h, Count: c, OriginalSet: o, UserSet: u}
	f.CreatedAt = time.Now()

	return f
}

func (f FunUntilHitDto) BsonD() bson.D {
	return bson.D{
		{"hits", f.Hits},
		{"count", f.Count},
		{"original", f.OriginalSet},
		{"custom", f.UserSet},
		{"createdAt", f.CreatedAt},
	}
}
