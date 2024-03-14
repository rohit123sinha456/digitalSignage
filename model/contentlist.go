package DataModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContentBlock struct {
	Type        string `bson:"contenttype,omitempty"`
	Content     string `bson:"contentid,omitempty"`
	DisplayTime int    `bson:"contentdisplaytime,omitempty"`
}

type ContentList struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	ContentList []ContentBlock     `bson:"contentlist"`
}
