package DTO

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FunSystemDataDto struct {
	Id          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	Hits        int                `bson:"hits"`
	Size        int                `bson:"size"`
	OriginalSet []int              `bson:"original"`
	UserSet     []int              `bson:"custom"`
}

func NewFunSystemDataDto(h int, s int, o []int, u []int) FunSystemDataDto {
	f := FunSystemDataDto{Hits: h, Size: s, OriginalSet: o, UserSet: u}
	f.CreatedAt = time.Now()

	return f
}

func (f FunSystemDataDto) BsonD() bson.D {
	return bson.D{
		{"hits", f.Hits},
		{"size", f.Size},
		{"original", f.OriginalSet},
		{"custom", f.UserSet},
		{"createdAt", f.CreatedAt},
	}
}
