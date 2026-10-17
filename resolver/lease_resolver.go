package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/vn-contrib/urhome-server/model"
)

type leaseResolver struct {
	model.Lease
}

func (r *leaseResolver) StartTime() graphql.Time {
	return resolveTime(r.Lease.StartTime)
}

func (r *leaseResolver) EndTime() *graphql.Time {
	return resolveNullTime(r.Lease.EndTime)
}

func (r *leaseResolver) Tenants() (trs []*tenantResolver, err error) {
	return
}
