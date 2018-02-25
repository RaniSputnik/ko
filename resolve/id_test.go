package resolve_test

import (
	"testing"

	"github.com/RaniSputnik/ko/resolve"
	graphql "github.com/neelance/graphql-go"
)

func TestEncodeID(t *testing.T) {
	testCases := []struct {
		Type, RawID string
		Expected    string
	}{
		{
			Type:     "Match",
			RawID:    "1",
			Expected: "TWF0Y2g6MQ==",
		},
		{
			Type:     "User",
			RawID:    "123456789",
			Expected: "VXNlcjoxMjM0NTY3ODk=",
		},
		{
			Type:     "Player",
			RawID:    "3dbb5f67-b662-4056-9b78-c029ffde3999",
			Expected: "UGxheWVyOjNkYmI1ZjY3LWI2NjItNDA1Ni05Yjc4LWMwMjlmZmRlMzk5OQ==",
		},
	}

	for _, testCase := range testCases {
		got := resolve.EncodeID(testCase.Type, testCase.RawID)
		if got != graphql.ID(testCase.Expected) {
			t.Errorf("Type='%s', RawID='%s', Expected='%s', Got='%s'",
				testCase.Type, testCase.RawID, testCase.Expected, got)
		}
	}
}
