package resolver

func New() *schemaResolver {
	return &schemaResolver{
		query: &queryResolver{},
	}
}

type schemaResolver struct {
	query *queryResolver
}

func (r *schemaResolver) Query() *queryResolver {
	return r.query
}
