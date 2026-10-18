package resolver

import (
	"context"

	"github.com/vn-contrib/urhome-server/model"
)

type tenantResolver struct {
	model.Tenant
}

func (r *tenantResolver) Profile(ctx context.Context) (*profileResolver, error) {
	p, err := r.Tenant.Profile().First(ctx)
	if err != nil {
		return nil, err
	}
	return &profileResolver{*p}, nil
}
