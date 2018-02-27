package main_test

import (
	"testing"

	"github.com/RaniSputnik/ko/resolve"
	"github.com/RaniSputnik/ko/schema"
	"github.com/neelance/graphql-go"
)

func TestParseSchema(t *testing.T) {
	if _, err := graphql.ParseSchema(schema.Text, resolve.Root(resolve.Data{})); err != nil {
		t.Error(err)
	}
}
