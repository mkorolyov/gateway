package main

import (
	"context"
	"github.com/99designs/gqlgen/handler"
	"github.com/mkorolyov/core/config"
	"github.com/mkorolyov/core/server"
	"github.com/mkorolyov/posts"
	profile "github.com/mkorolyov/profiles"
	"humans.net/ms/gateway"
)

const (
	defaultPort = "8080"
	profilePort = "9090"
	postsPort   = "9091"

	profilesTarget = "consul://0.0.0.0:8500/profiles"
	postsTarget    = "consul://0.0.0.0:8500/posts"
)

func main() {

	loader := config.Configure()
	s := server.New(loader)

	postsClient := posts.NewPostsClient(s.Connect("posts"))
	profilesClient := profile.NewProfileClient(s.Connect("profiles"))

	s.HandleHTTP("/query_playground", handler.Playground("GraphQL playground", "/query"))
	s.HandleHTTP("/query", handler.GraphQL(gateway.NewExecutableSchema(gateway.Config{
		Resolvers: gateway.NewResolver(postsClient, profilesClient)})))

	s.Serve(context.Background())
}
