package model

import (
	"github.com/uptrace/bun"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

var ModelLeaseTenant = bar.ModelFor[LeaseTenant]()

type LeaseTenant struct {
	bun.BaseModel `bun:"alias:lt"`

	LeaseID  int32 `bun:",pk"`
	TenantID int32 `bun:",pk"`
}
