package app

import(
	"context"
)

type contextKey struct {
	name string
}

type User struct {
	Name string
	IsAdmin bool
}

var userCtxKey = &contextKey{"user"}

func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}