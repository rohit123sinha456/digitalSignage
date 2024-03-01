package common

import "strings"

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
