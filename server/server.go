package main

import (
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

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gateway.NewExecutableSchema(gateway.Config{Resolvers: gateway.NewResolver(
		postsPort, profilePort)})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
