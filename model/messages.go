package DataModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type ImageBlock struct {
	Imagetype   string `bson:"imagetype,omitempty"`
	Image       string `bson:"image,omitempty"`
	DisplayTime int    `bson:"displaytime,omitempty"`
}

type DisplayBlock struct {
	BlockName string       `bson:"blockname,omitempty"`
	Imagelist []ImageBlock `bson:"imagelist,omitempty"`
}

type Playlist struct {
	ID           primitive.ObjectID `bson:"_id"`
	DeviceId     primitive.ObjectID `bson:"deviceid,omitempty"`
	DisplayBlock []DisplayBlock     `bson:"displayblock,omitempty"`
}
