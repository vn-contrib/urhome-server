package model

import (
	"github.com/uptrace/bun"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

type PropertyMemberRole string

const (
	PropertyMemberRoleOwner   PropertyMemberRole = "owner"
	PropertyMemberRoleManager PropertyMemberRole = "manager"
	PropertyMemberRoleTenant  PropertyMemberRole = "tenant"
)

var ModelPropertyMember = bar.ModelFor[PropertyMember]()

type PropertyMember struct {
	bun.BaseModel `bun:"alias:pm"`

	PropertyID int32 `bun:",pk"`
	UserID     int32 `bun:",pk"`
	Role       PropertyMemberRole
}
