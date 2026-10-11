package contexts

import (
	"context"

	"firebase.google.com/go/auth"
)

type fbAuthCtxKey struct{}

func WithFirebaseAuth(ctx context.Context, auth *auth.Client) context.Context {
	return context.WithValue(ctx, fbAuthCtxKey{}, auth)
}

func UseFirebaseAuth(ctx context.Context) *auth.Client {
	return ctx.Value(fbAuthCtxKey{}).(*auth.Client)
}
