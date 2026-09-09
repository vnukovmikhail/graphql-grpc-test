package main

import (
	"log"
	"net/http"

	"GRPC/tooldev/internal/graph"
	"GRPC/tooldev/internal/user"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	repo := user.NewMemoryRepository()
	svc := user.NewService(repo)

	resolver := &graph.Resolver{UserSvc: svc}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
