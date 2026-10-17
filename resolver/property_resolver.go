package resolver

import (
	"context"

	"github.com/vn-contrib/urhome-server/model"
)

type propertyResolver struct {
	model.Property
}

func (r *propertyResolver) Leases(ctx context.Context) (lrs []*leaseResolver, err error) {
	ls, err := r.Property.Leases().Take(ctx)
	for _, l := range ls {
		lrs = append(lrs, &leaseResolver{*l})
	}
	return lrs, err
}
