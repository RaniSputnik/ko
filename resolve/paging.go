package resolve

import graphql "github.com/neelance/graphql-go"

type pagingArgs struct {
	First  *int32
	After  *graphql.ID
	Last   *int32
	Before *graphql.ID
}
