package common

import (
    "path/filepath"
    "strings"
)

// System Identifier is how the system Identified a User across
func ExtractUserSystemIdentifier(userID string) string {
	useridsplits := strings.Split(userID, "-")
	userdBname := strings.Join([]string{"DSU", useridsplits[0]}, "")
	return userdBname

}

// vhost name for the rabbitmq
func CreatevHostName(userID string) string {
	useridsplits := strings.Split(userID, "-")
	userdBname := strings.Join([]string{"DSU", "VHOST", useridsplits[0]}, "")
	return userdBname

}

//bucket name for minio
func CreateBucketName(userID string) string {
	useridsplits := strings.Split(userID, "-")
	userdBname := strings.Join([]string{"dsu", "bucket", useridsplits[0]}, "")
	return userdBname

}

func GetFileType(filename string) string {
    ext := strings.ToLower(filepath.Ext(filename))
    imageExtensions := map[string]bool{
        ".jpeg": true,
        ".jpg":  true,
        ".webp": true,
        ".png":  true,
    }
    videoExtensions := map[string]bool{
        ".mp4": true,
    }

    if imageExtensions[ext] {
        return "Image"
    }

    if videoExtensions[ext] {
        return "Video"
    }

	return "Unknown"
}