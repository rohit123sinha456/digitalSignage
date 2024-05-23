package DataModel

import "go.mongodb.org/mongo-driver/bson/primitive"

// type ScreenBlock struct {
// 	BlockName     string             `bson:"blockname,omitempty"`
// 	ContentListID primitive.ObjectID `bson:"contentlistid,omitempty"`
// }

type Screen struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Location string             `bson:"location"`
	// Screenblock []ScreenBlock      `bson:"screenblock"`
	Screenblock string `bson:"screenblock"`
	Screencode  string `bson:"screencode"`
}
