package graph

import (
	"errors"

	"github.com/Holdrulff/autocomplete/internal/autocomplete"
)

var ErrNilService = errors.New("autocomplete service must not be nil")

type Resolver struct {
	service *autocomplete.Service
}

func NewResolver(service *autocomplete.Service) (*Resolver, error) {
	if service == nil {
		return nil, ErrNilService
	}

	return &Resolver{
		service: service,
	}, nil
}
