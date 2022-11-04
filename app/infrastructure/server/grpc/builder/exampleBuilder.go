package builder

import (
	"github.com/evenyosua18/template-project-structure/app/infrastructure/proto/pb"
	"github.com/mitchellh/mapstructure"
)

type ExampleBuilder struct{}

func (ExampleBuilder) GetExampleResponse(in interface{}) (interface{}, error) {
	var res *pb.ExampleResponse
	if err := mapstructure.Decode(in, &res); err != nil {
		return nil, err
	}

	return res, nil
}
