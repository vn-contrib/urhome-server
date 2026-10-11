package contexts

import (
	"context"
	"sync"

	"github.com/vn-contrib/urhome-server/model"
)

type authCtxKey struct{}

func WithAuth(ctx context.Context, fn func() *model.User) context.Context {
	return context.WithValue(ctx, authCtxKey{}, &authContext{
		User: sync.OnceValue(fn),
	})
}

func UseAuth(ctx context.Context) *authContext {
	return ctx.Value(authCtxKey{}).(*authContext)
}

type authContext struct {
	User func() *model.User
}
