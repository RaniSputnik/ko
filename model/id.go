package model

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Kind string

const (
	KindUnknown = Kind("")
	KindMatch   = Kind("Match")
	KindUser    = Kind("User")
)

func EncodeID(kind Kind, rawID string) string {
	bytes := []byte(fmt.Sprintf("%s:%s", kind, rawID))
	encodedID := base64.StdEncoding.EncodeToString(bytes)
	return encodedID
}

func DecodeID(encodedID string) (Kind, string) {
	bytes, err := base64.StdEncoding.DecodeString(encodedID)
	if err != nil {
		return KindUnknown, ""
	}
	parts := strings.Split(string(bytes), ":")
	if len(parts) != 2 {
		return KindUnknown, ""
	}
	return Kind(parts[0]), parts[1]
}
