package DataModel
// import "go.mongodb.org/mongo-driver/bson/primitive"

type EventStreamRequest struct {
	Screencode string               `bson:"screencode"`
	Userinfo   UserSystemIdentifeir `bson:"userinfo"`
	ScreenMongoID string `bson:"screenmongoid"`
}

type ScreenControlRequest struct {
	SignalType string               `bson:"signaltype"`
	Signal   UserSystemIdentifeir `bson:"signal"`
	ScreenMongoID string `bson:"screenmongoid"`
}
