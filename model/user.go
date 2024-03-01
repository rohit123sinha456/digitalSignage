package DataModel

type User struct {
	Name    string   `bson:"name,omitempty"`
	UserID  string   `bson:"user_id,omitempty"`
	Devices []string `bson:"devices,omitempty"`
}
