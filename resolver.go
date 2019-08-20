//go:generate gorunpkg github.com/99designs/gqlgen

package gateway

import (
	"context"
	"fmt"
	"github.com/mkorolyov/posts"
	profile "github.com/mkorolyov/profiles"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Resolver struct {
	postsClient    posts.PostsClient
	profilesClient profile.ProfileClient
}

func NewResolver(postsPort, profilePort string) *Resolver {
	connOpts := []grpc.DialOption{grpc.WithInsecure()}
	postsConn, err := grpc.Dial(":"+postsPort, connOpts...)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to grpc posts :%s: %v", postsPort, err))
	}

	postsClient := posts.NewPostsClient(postsConn)

	profileConn, err := grpc.Dial(":"+profilePort, connOpts...)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to grpc profile :%s: %v", postsPort, err))
	}

	profilesClient := profile.NewProfileClient(profileConn)

	return &Resolver{
		postsClient:    postsClient,
		profilesClient: profilesClient,
	}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

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
