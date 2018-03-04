package kontext

import (
	"context"
)

var userKey = User{}

func GetUser(ctx context.Context) User {
	return ctx.Value(userKey).(User)
}
func WithUser(parent context.Context, user User) context.Context {
	return context.WithValue(parent, userKey, user)
}
