package router

import (
	"database/sql"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/iqthuc/sport-shop/internal/delivery/graph/dataloader"
	graph "github.com/iqthuc/sport-shop/internal/delivery/graph/generated"
	"github.com/iqthuc/sport-shop/internal/delivery/graph/resolver"
	"github.com/vektah/gqlparser/v2/ast"
)

func IntGraphqlRouter(r *http.ServeMux, db *sql.DB) {

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{
		DB: db,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	sv := dataloader.Middleware(db, srv)
	r.Handle("/graphql", playground.Handler("GraphQL Playground", "/graphql/query"))
	r.Handle("/graphql/query", sv)
}
