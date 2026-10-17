package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go"
)

func New() *schemaResolver {
	return &schemaResolver{
		query:    &queryResolver{},
		mutation: &mutationResolver{},
	}
}

type schemaResolver struct {
	query    *queryResolver
	mutation *mutationResolver
}

func (r *schemaResolver) Query() *queryResolver {
	return r.query
}

func (r *schemaResolver) Mutation() *mutationResolver {
	return r.mutation
}

func resolveTime(t time.Time) graphql.Time {
	return graphql.Time{Time: t}
}

func resolveNullTime(t *time.Time) *graphql.Time {
	if t == nil {
		return nil
	}
	return &graphql.Time{Time: *t}
}
