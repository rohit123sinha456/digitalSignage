package DataModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Content struct {
	ID    primitive.ObjectID `bson:"_id"`
	DType string             `bson:"dtype"`
	Link  string             `bson:"link"`
}
