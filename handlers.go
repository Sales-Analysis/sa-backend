package main

import (
	"net/http"
	"sa-back/graph"
	"sa-back/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func gqlHandler() *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{}},
		),
	)
}

func handlers(mux *http.ServeMux, fileStaticHandler http.Handler) {
	mux.Handle("/", http.StripPrefix("/static/", fileStaticHandler))
	mux.Handle("/schema", playground.Handler("GraphQL playground", "/graphql"))
	mux.Handle("/graphql", gqlHandler())
	mux.HandleFunc("/upload", uploadHandler)
}
