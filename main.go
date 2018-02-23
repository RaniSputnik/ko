package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RaniSputnik/ko/resolve"
	"github.com/RaniSputnik/ko/schema"
	"github.com/neelance/graphql-go"

	"github.com/RaniSputnik/ko/handle"
)

func main() {
	s := graphql.MustParseSchema(schema.Text, resolve.Root())

	http.Handle("/", handle.GraphiQL("Ko", "/graphql"))
	http.Handle("/graphql", handle.GraphQL(s))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
