package model

import (
	"github.com/uptrace/bun"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

var ModelTenant = bar.ModelFor[Tenant]()

type Tenant struct {
	bun.BaseModel `bun:"alias:t"`

	ID   int32 `bun:",pk,autoincrement"`
	Name string
}
