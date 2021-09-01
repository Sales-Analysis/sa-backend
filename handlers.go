package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
	"sa-back/graph"
	"sa-back/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
)

func gqlHandler() *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{}},
		),
	)
}

func handlers(mux *http.ServeMux, fileStaticHandler http.Handler) {
	mux.Handle("/", http.StripPrefix("/data/", fileStaticHandler))
	mux.Handle("/query", playground.Handler("GraphQL playground", "/graphql"))
	mux.Handle("/graphql", gqlHandler())
	mux.HandleFunc("/upload", uploadHandler)
}
