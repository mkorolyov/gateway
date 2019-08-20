package main

import (
	"fmt"
	"github.com/mkorolyov/posts"
	profile "github.com/mkorolyov/profiles"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"humans.net/ms/gateway"
)

const (
	defaultPort = "8080"
	profilePort = "9090"
	postsPort   = "9091"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	connOpts := []grpc.DialOption{grpc.WithInsecure()}
	postsConn, err := grpc.Dial("0.0.0.0:9091", connOpts...)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to grpc posts :%s: %v", postsPort, err))
	}

	postsClient := posts.NewPostsClient(postsConn)

	profileConn, err := grpc.Dial("0.0.0.0:9090", connOpts...)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to grpc profile :%s: %v", postsPort, err))
	}

	profilesClient := profile.NewProfileClient(profileConn)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query",
		handler.GraphQL(gateway.NewExecutableSchema(gateway.Config{
			Resolvers: gateway.NewResolver(postsClient, profilesClient)})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
