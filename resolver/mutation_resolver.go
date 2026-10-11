package resolver

import (
	"context"

	"github.com/vn-contrib/urhome-server/contexts"
	"github.com/vn-contrib/urhome-server/model"
	"github.com/vn-contrib/urhome-server/resolver/input"
)

type mutationResolver struct{}

func (*mutationResolver) CreateProperty(ctx context.Context, args struct {
	Input input.PropertyInput
}) (*propertyResolver, error) {
	p := new(model.Property)
	if _, err := args.Input.Map(p); err != nil {
		return nil, err
	}
	err := contexts.UseAuth(ctx).User().CreateProperty(ctx, p)
	return &propertyResolver{*p}, err

}
