package bar

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

func ModelFor[T any]() Model[T] {
	return Model[T]{}
}

type Model[T any] struct {
	fields []Field
}

func (m Model[T]) Where(fields ...Field) Model[T] {
	return Model[T]{
		fields: append(m.fields, fields...),
	}
}

func (m Model[T]) Create(ctx context.Context, fields ...Field) (*T, error) {
	model, fns := new(T), m.Where(fields...).insertQueries()
	_, err := DBFromContext(ctx).NewInsert().Model(model).Apply(fns...).Exec(ctx)
	return model, err
}

func (m Model[T]) First(ctx context.Context) (*T, error) {
	model, fns := new(T), m.selectQueries()
	err := DBFromContext(ctx).NewSelect().Model(model).Apply(fns...).Limit(1).Scan(ctx)
	return model, err
}

func (m Model[T]) FirstOrCreate(ctx context.Context) (*T, error) {
	model, err := m.First(ctx)
	if err == sql.ErrNoRows {
		return m.Create(ctx)
	}
	return model, err
}

func (m Model[T]) selectQueries() (fns []func(*bun.SelectQuery) *bun.SelectQuery) {
	for _, val := range m.fields {
		fns = append(fns, val.SelectQuery)
	}
	return
}

func (m Model[T]) insertQueries() (fns []func(*bun.InsertQuery) *bun.InsertQuery) {
	for _, val := range m.fields {
		fns = append(fns, val.InsertQuery)
	}
	return
}
