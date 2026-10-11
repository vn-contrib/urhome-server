package input

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/vn-contrib/urhome-server/model"
)

type PropertyInput struct {
	Name graphql.NullString
}

func (i PropertyInput) Map(p *model.Property) (cols []string, err error) {
	if i.Name.Set {
		p.Name = *i.Name.Value
		cols = append(cols, "name")
	}
	return
}
