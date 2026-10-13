package model

import (
	"github.com/uptrace/bun"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

var ModelUser = bar.ModelFor[User]()

type User struct {
	bun.BaseModel `bun:"alias:u"`

	ID          int32 `bun:",pk,autoincrement"`
	FirebaseUID string
}

func (u User) PropertyMembers() bar.Model[PropertyMember] {
	return ModelPropertyMember.Where(bar.Field{"user_id": u.ID})
}
