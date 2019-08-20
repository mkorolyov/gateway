package gateway

import (
	"context"
	"github.com/mkorolyov/posts"
	profile "github.com/mkorolyov/profiles"
	"github.com/pkg/errors"
)

type ModificationResolver struct {
	postsClient    posts.PostsClient
	profilesClient profile.ProfileClient
}

func (mr ModificationResolver) CreateProfile(ctx context.Context, firstName string, lastName string) (string, error) {
	res, err := mr.profilesClient.Create(ctx, &profile.CreateRequest{
		FirstName: firstName,
		LastName:  lastName,
	})
	if err != nil {
		return "", errors.Wrapf(err, "failed to create profile for %s %s", firstName, lastName)
	}

	return res.Id, nil
}

func (mr ModificationResolver) PublishPost(ctx context.Context, userID string, name string, description string, typeArg string) (string, error) {
	res, err := mr.postsClient.Create(ctx, &posts.CreateRequest{
		UserId:      userID,
		Name:        name,
		Description: description,
		Type:        typeArg,
	})

	if err != nil {
		return "", errors.Wrapf(err, "failed to create post %s for user %s", name, userID)
	}

	return res.Id, nil
}

func New(postsClient posts.PostsClient, profilesClient profile.ProfileClient) *ModificationResolver {
	return &ModificationResolver{
		postsClient:    postsClient,
		profilesClient: profilesClient,
	}
}
