package DataModel

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type ScreenBlock struct {
	BlockName     string             `bson:"blockname,omitempty"`
	ContentListID primitive.ObjectID `bson:"contentlistid,omitempty"`
}

type Screen struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Location string             `bson:"location"`
	CurrentPlaylistName string  `bson:"currentplaylistname,omitempty"`
	CurrentPlaylistID primitive.ObjectID    `bson:"currentplaylistid,omitempty"`
	// PlaylistList []primitive.ObjectID    `bson:"playlistlistid,omitempty"`
	// Screenblock []ScreenBlock      `bson:"screenblock"`
	Screenblock string `bson:"screenblock"`
	Screencode  string `bson:"screencode"`
	CreatedAt *time.Time         `bson:"createdAt,omitempty"`
}
