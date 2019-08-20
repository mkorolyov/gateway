//go:generate gorunpkg github.com/99designs/gqlgen

package gateway

import context "context"

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Profile(ctx context.Context, id string) (*Profile, error) {
	panic("not implemented")
}
