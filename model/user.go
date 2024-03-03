package DataModel

type User struct {
	Name    string   `bson:"name,omitempty"`
	UserID  string   `bson:"userid,omitempty"`
	Devices []Device `bson:"devices,omitempty"`
}

type UserSystemIdentifeir struct {
	UserBucketID     string `bson:"userbucketid,omitempty"`
	UserQueueID      string `bson:"userqueueid,omitempty"`
	UserQueuevHostID string `bson:"userqueuevhostid,omitempty"`
	UserSystemID     string `bson:"usersystemid,omitempty"`
	UserID           string `bson:"userid,omitempty"`
}

type Device struct {
	DID  string `bson:"name,omitempty"`
	Name string `bson:"name,omitempty"`
}
