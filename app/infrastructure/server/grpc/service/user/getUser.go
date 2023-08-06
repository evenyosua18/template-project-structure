package user

import (
	"context"
	"github.com/evenyosua18/template-project-structure/app/infrastructure/proto/pb"
)

func (s *ServiceUser) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	//call interaction
	res, err := s.uc.GetUser(ctx, in)
	if err != nil {
		return nil, err
	}

	return res.(*pb.UserResponse), nil
}
