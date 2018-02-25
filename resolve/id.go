package resolve

import (
	"encoding/base64"
	"fmt"

	graphql "github.com/neelance/graphql-go"
)

const (
	matchID = "Match"
)

func EncodeID(kind, rawid string) graphql.ID {
	bytes := []byte(fmt.Sprintf("%s:%s", kind, rawid))
	encodedID := base64.StdEncoding.EncodeToString(bytes)
	return graphql.ID(encodedID)
}
