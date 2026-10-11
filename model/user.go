package model

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

type User struct {
	bun.BaseModel `bun:"alias:u"`

	ID          int32 `bun:",pk,autoincrement"`
	FirebaseUID string
}

func (u User) CreateProperty(ctx context.Context, p *Property) error {
	return bar.RunInTx(ctx, nil, func(ctx context.Context) error {
		if err := bar.Create(ctx, p); err != nil {
			return err
		}

		if err := bar.Create(ctx, &PropertyMember{
			PropertyID: p.ID,
			UserID:     u.ID,
			Role:       PropertyMemberRoleOwner,
		}); err != nil {
			return err
		}

		return nil
	})
}
