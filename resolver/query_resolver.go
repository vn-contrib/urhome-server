package resolver

import "context"

type queryResolver struct{}

func (*queryResolver) Ping(ctx context.Context) string {
	return "pong"
}
