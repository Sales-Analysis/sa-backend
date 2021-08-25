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
	mux.Handle("/graphql", playground.Handler("GraphQL playground", "/schema"))
	mux.Handle("/schema", gqlHandler())
	mux.HandleFunc("/upload", uploadHandler)
}
