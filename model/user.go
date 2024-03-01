package DataModel

type User struct {
	Name    string   `bson:"name,omitempty"`
	UserID  string   `bson:"user_id,omitempty"`
	Devices []Device `bson:"devices,omitempty"`
}

type Device struct {
	DID  string `bson:"name,omitempty"`
	Name string `bson:"name,omitempty"`
}
