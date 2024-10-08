package DataModel

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Content struct {
	ID    primitive.ObjectID `bson:"_id"`
	CName string             `bson:"cname"`
	DType string             `bson:"dtype"`
	Link  string             `bson:"link"`
	CreatedAt *time.Time         `bson:"createdAt,omitempty"`
}
