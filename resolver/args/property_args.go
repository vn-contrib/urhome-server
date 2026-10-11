package args

import (
	"context"

	"github.com/dundunlabs/grapher/scalar"
	"github.com/vn-contrib/urhome-server/contexts"
	"github.com/vn-contrib/urhome-server/model"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

type PropertyArgs struct {
	PropertyID scalar.IntID
}

func (args PropertyArgs) Member(ctx context.Context) *model.PropertyMember {
	f := bar.Field{
		"user_id":     contexts.UseAuth(ctx).User().ID,
		"property_id": args.PropertyID,
	}
	pm, err := model.ModelPropertyMember.Where(f).First(ctx)
	if err != nil {
		panic(err)
	}
	return pm
}
