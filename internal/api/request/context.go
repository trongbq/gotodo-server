package request

import (
	"context"

	"github.com/trongbq/gotodo-server/internal"
)

type key int

const userKey key = iota

func WithUser(origin context.Context, user *internal.User) context.Context {
	return context.WithValue(origin, userKey, user)
}

func UserFrom(ctx context.Context) (*internal.User, bool) {
	user, ok := ctx.Value(userKey).(*internal.User)
	return user, ok
}
