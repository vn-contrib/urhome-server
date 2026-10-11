package model

import "github.com/uptrace/bun"

type Property struct {
	bun.BaseModel `bun:"alias:p"`

	ID   int32 `bun:",pk,autoincrement"`
	Name string
}
