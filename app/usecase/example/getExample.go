package example

import (
	"context"
	"github.com/evenyosua18/template-project-structure/app/domain/entity"
	"github.com/mitchellh/mapstructure"
)

func (i *InteractionExample) GetExample(ctx context.Context, in interface{}) (interface{}, error) {
	//decode
	var request *entity.ExampleEntityRequest
	if err := mapstructure.Decode(in, &request); err != nil {
		return nil, err
	}

	//call repository
	if err := i.repo.GetExample(ctx, in); err != nil {
		return nil, err
	}

	return i.out.GetExampleResponse(&entity.ExampleEntityResponse{Message: request.Message})
}
