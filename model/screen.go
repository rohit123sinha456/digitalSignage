package DataModel

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type ScreenBlock struct {
	BlockName     string             `bson:"blockname,omitempty"`
	ContentListID primitive.ObjectID `bson:"contentlistid,omitempty"`
}

type PlaylistsofScreen struct {
	PlaylistName     string             `bson:"playlistname,omitempty"`
	PlaylistID primitive.ObjectID `bson:"_id,omitempty"`
}

type Screen struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name,omitempty"`
	Location string             `bson:"location,omitempty"`
	CurrentPlaylistName string  `bson:"currentplaylistname,omitempty"`
	CurrentPlaylistID primitive.ObjectID    `bson:"currentplaylistid,omitempty"`
	CreatedAt *time.Time         `bson:"createdAt,omitempty"`
	UpdatedAt *time.Time         `bson:"updatedAt,omitempty"`
	Status bool `bson:"status,omitempty"`
	Orientation int `bson:"orientation,omitempty"`
	StorageTotal float32 `bson:"storagetotal,omitempty"`
	StorageFree float32 `bson:"storagefree,omitempty"`
	StorageUsed float32 `bson:"storageused,omitempty"`
	MemoryTotal float32 `bson:"memorytotal,omitempty"`
	MemoryUsed float32 `bson:"memoryused,omitempty"`
	IPAddr string `bson:"ip,omitempty"`
	DeviceModel string `bson:"devicemodel,omitempty"`
	CanDrawOverlay bool `bson:"candrawoverlay,omitempty"`
	AppShellVersion string `bson:"appshellversion,omitempty"`
	ScreenshotSupport bool `bson:"screenshotsupport,omitempty"`
	ScreenResolution string `bson:"screenresolution,omitempty"`
	BrowserResolution string `bson:"browserresolution,omitempty"`
	EngerySavedEnabled bool `bson:"energysaver,omitempty"`
	Country string `bson:"country,omitempty"`
	UserAgent string `bson:"useragent,omitempty"`
	GPlaySupport bool `bson:"gplaysupport,omitempty"`
	VideoCodecs string `bson:"videocodecs,omitempty"`
	PlayerTimezone string `bson:"playertimsezone,omitempty"`
	OS string `bson:"os,omitempty"`
	DevicePixelRatio string `bson:"devicepixelratio,omitempty"`
	PlayerCodec string `bson:"playercodec,omitempty"`
	RicoviAppVersion string `bson:"appversion,omitempty"`
	// PlaylistList []primitive.ObjectID    `bson:"playlistlistid,omitempty"`
	// Screenblock []ScreenBlock      `bson:"screenblock"`
}
