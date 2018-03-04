package model_test

import (
	"testing"

	"github.com/RaniSputnik/ko/model"
)

func TestEncodeID(t *testing.T) {
	testCases := []struct {
		Kind     model.Kind
		RawID    string
		Expected string
	}{
		{
			Kind:     model.KindMatch,
			RawID:    "1",
			Expected: "TWF0Y2g6MQ==",
		},
		{
			Kind:     model.KindUser,
			RawID:    "123456789",
			Expected: "VXNlcjoxMjM0NTY3ODk=",
		},
		{
			Kind:     model.Kind("Player"),
			RawID:    "3dbb5f67-b662-4056-9b78-c029ffde3999",
			Expected: "UGxheWVyOjNkYmI1ZjY3LWI2NjItNDA1Ni05Yjc4LWMwMjlmZmRlMzk5OQ==",
		},
	}

	for _, testCase := range testCases {
		got := model.EncodeID(testCase.Kind, testCase.RawID)
		if got != testCase.Expected {
			t.Errorf("Kind='%s', RawID='%s', Expected='%s', Got='%s'",
				testCase.Kind, testCase.RawID, testCase.Expected, got)
		}
	}
}

func TestDecodeID(t *testing.T) {
	testCases := []struct {
		ID         string
		ExpectKind model.Kind
		ExpectID   string
	}{
		// Success cases
		{
			ID:         "VXNlcjoxMjM0NQ==",
			ExpectKind: model.Kind("User"),
			ExpectID:   "12345",
		},
		{
			ID:         "TWF0Y2g6YXNkamJqaGJyb2Vy",
			ExpectKind: model.KindMatch,
			ExpectID:   "asdjbjhbroer",
		},
		{
			ID:         "Zm9vIGJhcjpramFuc2Q9YTA5dTEyM24=",
			ExpectKind: model.Kind("foo bar"),
			ExpectID:   "kjansd=a09u123n",
		},
		{
			ID:         "VXNlcjoxMjM0NTY3ODk=",
			ExpectKind: model.KindUser,
			ExpectID:   "123456789",
		},
		// Failure cases
		{
			ID:         "foo bar",
			ExpectKind: model.KindUnknown,
			ExpectID:   "",
		},
		{
			ID:         "Player:12345",
			ExpectKind: model.KindUnknown,
			ExpectID:   "",
		},
		{
			ID:         "UGxheWVy",
			ExpectKind: model.KindUnknown,
			ExpectID:   "",
		},
	}

	for _, testCase := range testCases {
		gotKind, gotID := model.DecodeID(testCase.ID)
		if gotKind != testCase.ExpectKind {
			t.Errorf("ID='%s', Expected kind='%s', Got='%s'",
				testCase.ID, testCase.ExpectKind, gotKind)
		}
		if gotID != testCase.ExpectID {
			t.Errorf("ID='%s', Expected id='%s', Got='%s'",
				testCase.ID, testCase.ExpectID, gotID)
		}
	}
}
