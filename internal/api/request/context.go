package request

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/trongbq/gotodo-server/internal"
)

type key int

const (
	userKey key = iota
	logKey
)

func WithUser(origin context.Context, user *internal.User) context.Context {
	return context.WithValue(origin, userKey, user)
}

func UserFrom(ctx context.Context) (*internal.User, bool) {
	user, ok := ctx.Value(userKey).(*internal.User)
	return user, ok
}

func WithLog(origin context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(origin, logKey, logger)
}

func LogFrom(ctx context.Context) (*logrus.Entry, bool) {
	logger, ok := ctx.Value(logKey).(*logrus.Entry)
	return logger, ok
}
