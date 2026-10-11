package input

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

type PropertyInput struct {
	Name graphql.NullString
}

func (i PropertyInput) Field() (bar.Field, error) {
	f := bar.Field{}
	if i.Name.Set {
		f["name"] = *i.Name.Value
	}
	return f, nil
}
