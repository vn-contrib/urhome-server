package bar

import "github.com/uptrace/bun"

type Field map[string]any

func (f Field) SelectQuery(sq *bun.SelectQuery) *bun.SelectQuery {
	for k, v := range f {
		sq = sq.Where("? = ?", bun.Ident(k), v)
	}
	return sq
}

func (f Field) InsertQuery(iq *bun.InsertQuery) *bun.InsertQuery {
	for k, v := range f {
		iq = iq.Value(k, "?", v)
	}
	return iq
}

func (f Field) UpdateQuery(uq *bun.UpdateQuery) *bun.UpdateQuery {
	for k, v := range f {
		uq = uq.Value(k, "?", v)
	}
	return uq
}
