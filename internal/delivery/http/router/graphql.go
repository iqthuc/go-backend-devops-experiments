package router

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/iqthuc/sport-shop/internal/delivery/graph"
	"github.com/vektah/gqlparser/v2/ast"
)

func IntGraphqlRouter(r *http.ServeMux) {

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	// authRouter.HandleFunc("GET /profile", middleware.Wrap(authHandler.GetUserInfo, authMW.Auth))
	r.Handle("/graphql", playground.Handler("GraphQL Playground", "/graphql/query"))
	r.Handle("/graphql/query", srv)
}
