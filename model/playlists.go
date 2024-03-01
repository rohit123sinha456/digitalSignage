package DataModel

type PlaylistType int

const (
	Single PlaylistType = iota
	Multiple
)

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
	ID           string         `bson:"id,omitempty"`
	DeviceId     string         `bson:"deviceid,omitempty"`
	PType        PlaylistType   `bson:"ptype,omitempty"`
	DisplayBlock []DisplayBlock `bson:"displayblock,omitempty"`
}
