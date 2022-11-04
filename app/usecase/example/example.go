package example

import (
	"context"
	"github.com/evenyosua18/template-project-structure/app/domain/repository"
)

type InputPortExample interface {
	GetExample(ctx context.Context, in interface{}) (interface{}, error)
}

type OutputPortExample interface {
	GetExampleResponse(in interface{}) (interface{}, error)
}

type InteractionExample struct {
	repo repository.ExampleRepository
	out  OutputPortExample
}

func NewInteractionExample(r repository.ExampleRepository, o OutputPortExample) *InteractionExample {
	return &InteractionExample{
		repo: r,
		out:  o,
	}
}
