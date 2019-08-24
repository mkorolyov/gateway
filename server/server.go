package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/mkorolyov/core/discovery/consul"
	"github.com/mkorolyov/posts"
	profile "github.com/mkorolyov/profiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
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
	consul.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	connOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name)}

	postsConsulConn, err := grpc.DialContext(ctx, postsTarget, connOpts...)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to grpc posts: %v", err))
	}
	defer postsConsulConn.Close()
	postsClient := posts.NewPostsClient(postsConsulConn)

	profilesConsulConn, err := grpc.DialContext(ctx, profilesTarget, connOpts...)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to grpc profile: %v", err))
	}
	defer profilesConsulConn.Close()
	profilesClient := profile.NewProfileClient(profilesConsulConn)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query",
		handler.GraphQL(gateway.NewExecutableSchema(gateway.Config{
			Resolvers: gateway.NewResolver(postsClient, profilesClient)})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
