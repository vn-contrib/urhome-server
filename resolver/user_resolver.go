package resolver

import (
	"context"

	"github.com/vn-contrib/urhome-server/model"
)

type userResolver struct {
	model.User
}

func (r *userResolver) Properties(ctx context.Context) (rs []*propertyMemberResolver, err error) {
	pms, err := r.PropertyMembers().Take(ctx)
	if err != nil {
		return nil, err
	}

	for _, pm := range pms {
		rs = append(rs, &propertyMemberResolver{*pm})
	}
	return
}
