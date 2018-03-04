package handle

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/RaniSputnik/ko/kontext"
	"github.com/RaniSputnik/ko/model"
	jwt "github.com/dgrijalva/jwt-go"
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

		ctx, gotUser := getUserContext(r)
		if !gotUser {
			unauthorized(w)
			return
		}

		res := schema.Exec(ctx, p.Query, p.OperationName, p.Variables)
		ok(w, res)
	}
}

const bearerPrefix = "Bearer "

func getUserContext(r *http.Request) (context.Context, bool) {
	userContextInvalid := r.Context()

	authorization := r.Header.Get("Authorization")
	if !strings.HasPrefix(authorization, bearerPrefix) {
		return userContextInvalid, false
	}

	tokenString := strings.TrimPrefix(authorization, bearerPrefix)
	token, err := jwt.Parse(tokenString, todoVerifyToken)
	if err != nil {
		return userContextInvalid, false
	}
	claims := token.Claims.(jwt.MapClaims)
	kind, id := model.DecodeID(claims["sub"].(string))
	if kind != model.KindUser {
		return userContextInvalid, false
	}

	return kontext.WithUser(r.Context(), model.User{ID: id}), true
}

func todoVerifyToken(token *jwt.Token) (interface{}, error) {
	return jwt.UnsafeAllowNoneSignatureType, nil
}

func parseParams(r *http.Request) (params, error) {
	// TODO support GET method

	var p params
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	return p, err
}
