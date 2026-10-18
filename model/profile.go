package model

import (
	"github.com/uptrace/bun"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

var ModelProfile = bar.ModelFor[Profile]()

type Profile struct {
	bun.BaseModel `bun:"alias:pf"`

	ID     int32 `bun:",pk,autoincrement"`
	Name   string
	UserID *int32
}
