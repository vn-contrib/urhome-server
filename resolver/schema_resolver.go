package resolver

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
