package handle

import (
	"encoding/json"
	"net/http"

	"github.com/neelance/graphql-go"
)

type params struct {
	Query         string
	OperationName string
	Variables     map[string]interface{}
}

func GraphQL(schema *graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := parseParams(r)
		if err != nil {
			badRequest(w, err)
			return
		}

		res := schema.Exec(r.Context(), p.Query, p.OperationName, p.Variables)
		ok(w, res)
	}
}

func parseParams(r *http.Request) (params, error) {
	// TODO support GET method

	var p params
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	return p, err
}
