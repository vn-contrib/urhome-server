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
	pm := &model.PropertyMember{
		UserID:     contexts.UseAuth(ctx).User().ID,
		PropertyID: int32(args.PropertyID),
	}
	if err := bar.FindBy(ctx, pm); err != nil {
		panic(err)
	}
	return pm
}
