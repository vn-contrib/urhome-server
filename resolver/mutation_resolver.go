package resolver

import (
	"context"

	"github.com/vn-contrib/urhome-server/contexts"
	"github.com/vn-contrib/urhome-server/model"
	"github.com/vn-contrib/urhome-server/pkg/bar"
	"github.com/vn-contrib/urhome-server/resolver/input"
)

type mutationResolver struct{}

func (*mutationResolver) CreateProperty(ctx context.Context, args struct {
	Input input.PropertyInput
}) (pr *propertyResolver, err error) {
	f, err := args.Input.Field()
	if err != nil {
		return nil, err
	}

	err = bar.RunInTx(ctx, nil, func(ctx context.Context) error {
		p, err := model.ModelProperty.Create(ctx, f)
		if err != nil {
			return err
		}

		_, err = model.ModelPropertyMember.Create(ctx, bar.Field{
			"user_id":     contexts.UseAuth(ctx).User().ID,
			"property_id": p.ID,
			"role":        model.PropertyMemberRoleOwner,
		})
		if err != nil {
			return err
		}

		pr = &propertyResolver{*p}
		return nil
	})

	return
}
