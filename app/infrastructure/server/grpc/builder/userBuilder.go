package builder

import (
	"github.com/evenyosua18/template-project-structure/app/infrastructure/proto/pb"
	"github.com/evenyosua18/template-project-structure/app/usecase/user"
	"github.com/mitchellh/mapstructure"
)

type UserBuilder struct{}

func NewUserBuilder() user.IUserInteractionResponse {
	return &UserBuilder{}
}

func (UserBuilder) GetUserResponse(in interface{}) (interface{}, error) {
	var res *pb.UserResponse
	if err := mapstructure.Decode(in, &res); err != nil {
		return nil, err
	}

	return res, nil
}
