package model

import "github.com/uptrace/bun"

type PropertyMemberRole string

const (
	PropertyMemberRoleOwner   PropertyMemberRole = "owner"
	PropertyMemberRoleManager PropertyMemberRole = "manager"
	PropertyMemberRoleTenant  PropertyMemberRole = "tenant"
)

type PropertyMember struct {
	bun.BaseModel `bun:"alias:pm"`

	PropertyID int32 `bun:",pk"`
	UserID     int32 `bun:",pk"`
	Role       PropertyMemberRole
}
