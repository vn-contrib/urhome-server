package resolver

import (
	"context"

	"github.com/vn-contrib/urhome-server/model"
)

type propertyMemberResolver struct {
	model.PropertyMember
}

func (r *propertyMemberResolver) Property(ctx context.Context) (*propertyResolver, error) {
	p, err := r.PropertyMember.Property().First(ctx)
	if err != nil {
		return nil, err
	}
	return &propertyResolver{*p}, nil
}
