package DataModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type ImageBlock struct {
	ImageId     primitive.ObjectID `bson:"imageid,omitempty"`
	DisplayTime int                `bson:"displaytime,omitempty"`
}

type DisplayBlock struct {
	BlockName string       `bson:"blockname,omitempty"`
	Imagelist []ImageBlock `bson:"imagelist,omitempty"`
}

type DeviceBlock struct {
	DeviceId     primitive.ObjectID `bson:"deviceid,omitempty"`
	DisplayBlock []DisplayBlock     `bson:"displayblock,omitempty"`
}

type Playlist struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"playlistname,omitempty"`
	DeviceBlock []DeviceBlock      `bson:"deviceblock,omitempty"`
}

type UpdatePlaylistRequest struct {
	Name        string        `bson:"name,omitempty"`
	DeviceBlock []DeviceBlock `bson:"deviceblock,omitempty"`
}
