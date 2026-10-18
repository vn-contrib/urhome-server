package model

import (
	"github.com/uptrace/bun"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

var ModelTenant = bar.ModelFor[Tenant]()

type Tenant struct {
	bun.BaseModel `bun:"alias:t"`

	LeaseID   int32 `bun:",pk"`
	ProfileID int32 `bun:",pk"`
}

func (t Tenant) Profile() bar.Model[Profile] {
	return ModelProfile.Where(bar.Field{"id": t.ProfileID})
}
