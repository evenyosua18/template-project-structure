package example

import (
	"context"
	"github.com/evenyosua18/template-project-structure/app/infrastructure/proto/pb"
)

func (s *ServiceExample) GetExample(ctx context.Context, in *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	//call interaction
	res, err := s.uc.GetExample(ctx, in)
	if err != nil {
		return nil, err
	}

	return res.(*pb.ExampleResponse), nil
}
