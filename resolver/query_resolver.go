package resolver

import (
	"context"

	"github.com/vn-contrib/urhome-server/contexts"
	"github.com/vn-contrib/urhome-server/pkg/bar"
	"github.com/vn-contrib/urhome-server/resolver/args"
)

type queryResolver struct{}

func (*queryResolver) Ping(ctx context.Context) (string, error) {
	_, err := bar.DBFromContext(ctx).QueryContext(ctx, "SELECT 1")
	return "pong", err
}

func (*queryResolver) CurrentUser(ctx context.Context) *userResolver {
	u := contexts.UseAuth(ctx).User()
	return &userResolver{*u}
}

func (*queryResolver) CurrentPropertyMember(ctx context.Context, args struct {
	args.PropertyArgs
}) *propertyMemberResolver {
	pm := args.Member(ctx)
	return &propertyMemberResolver{*pm}
}
