//go:generate gorunpkg github.com/99designs/gqlgen

package gateway

import (
	"context"
	"github.com/mkorolyov/posts"
	profile "github.com/mkorolyov/profiles"
	"github.com/pkg/errors"
)

type Resolver struct {
	postsClient    posts.PostsClient
	profilesClient profile.ProfileClient
}

func NewResolver(postsClient posts.PostsClient, profileClient profile.ProfileClient) *Resolver {
	return &Resolver{
		postsClient:    postsClient,
		profilesClient: profileClient,
	}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() MutationResolver {
	return &modificationResolver{&ModificationResolver{
		postsClient:    r.postsClient,
		profilesClient: r.profilesClient,
	}}
}

type queryResolver struct{
	*Resolver
}

type modificationResolver struct {
	*ModificationResolver
}

func (r *queryResolver) Profile(ctx context.Context, id string) (*Profile, error) {
	p, err := r.profilesClient.Get(ctx, &profile.GetRequest{Id: id})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load profile %s", id)
	}

	ps, err := r.postsClient.Get(ctx, &posts.GetRequest{UserId: id})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load posts for profile %s", id)
	}

	res := &Profile{
		ID:        id,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Posts:     make([]*Post, 0, len(ps.Posts)),
	}

	for _, post := range ps.Posts {
		res.Posts = append(res.Posts, &Post{
			ID:          post.Id,
			Name:        post.Name,
			Description: post.Description,
			Type:        post.Type,
		})
	}

	return res, nil
}
