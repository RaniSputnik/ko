package model_test

import (
	"testing"

	"github.com/RaniSputnik/ko/model"
)

func TestMatchStatus(t *testing.T) {
	testCases := []struct {
		Match          model.Match
		ExpectedStatus model.MatchStatus
	}{
		{
			// Waiting for opponent
			Match:          model.Match{ID: "imwaiting", Owner: "Alice", BoardSize: 19},
			ExpectedStatus: model.MatchStatusWaiting,
		},
		{
			// Opponent has joined
			Match:          model.Match{ID: "ready", Owner: "Alice", BoardSize: 19, Opponent: "Bob"},
			ExpectedStatus: model.MatchStatusReady,
		},
		// TODO other statuses
	}

	for _, testCase := range testCases {
		if got := testCase.Match.Status(); got != testCase.ExpectedStatus {
			t.Errorf("Expected match '%v' to have status '%s' but got '%s'",
				testCase.Match, testCase.ExpectedStatus, got)
		}
	}
}
