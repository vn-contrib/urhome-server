package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"alias:u"`

	ID int32 `bun:",pk,autoincrement"`
}
