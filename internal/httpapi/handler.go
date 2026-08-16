package httpapi

import (
	"errors"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/Holdrulff/autocomplete/graph"
)

var ErrNilResolver = errors.New("graphql resolver must not be nil")

func NewHandler(resolver *graph.Resolver) (http.Handler, error) {
	if resolver == nil {
		return nil, ErrNilResolver
	}

	executableSchema := graph.NewExecutableSchema(
		graph.Config{
			Resolvers: resolver,
		},
	)

	return handler.NewDefaultServer(executableSchema), nil
}
