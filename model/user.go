package DataModel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name,omitempty"`
	UserID        string             `bson:"userid,omitempty"`
	Devices       []Device           `bson:"devices,omitempty"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"Password" validate:"required,min=6"`
	Email         *string            `json:"email" validate:"email,required"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}

type UserSystemIdentifeir struct {
	UserBucketID     string `bson:"userbucketid,omitempty"`
	UserQueueID      string `bson:"userqueueid,omitempty"`
	UserQueuevHostID string `bson:"userqueuevhostid,omitempty"`
	UserSystemID     string `bson:"usersystemid,omitempty"`
	UserID           string `bson:"userid,omitempty"`
}

type Device struct {
	DID  string `bson:"deviceid,omitempty"`
	Name string `bson:"devicename,omitempty"`
}
