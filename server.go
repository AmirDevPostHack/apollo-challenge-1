package main

import (
    "log"
    "net/http"
    "os"

    "github.com/amir/graphql-c1/internal/service"
    "github.com/amir/graphql-c1/pkg/auth"
    "github.com/go-chi/chi/v5"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"

    "github.com/amir/graphql-c1/graph"
)

const defaultPort = "8080"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    router := chi.NewRouter()
    router.Use(auth.Middleware)

    srv := handler.NewDefaultServer(
        graph.NewExecutableSchema(
            graph.Config{Resolvers: graph.NewResolver(service.New())},
        ),
    )

    router.Handle("/", playground.Handler("GraphQL playground", "/query"))
    router.Handle("/query", srv)

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
