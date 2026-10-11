package model

import (
	"github.com/uptrace/bun"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

var ModelProperty = bar.ModelFor[Property]()

type Property struct {
	bun.BaseModel `bun:"alias:p"`

	ID   int32 `bun:",pk,autoincrement"`
	Name string
}
