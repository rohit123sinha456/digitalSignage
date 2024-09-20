package DataModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Signal struct{
	SignalType     string  `bson:"signaltype,omitempty"`
	ScreenID string  `bson:"screenid,omitempty"`
	IDs []primitive.ObjectID `bson:"ids,omitempty"`
}