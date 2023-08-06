package user

import (
	"github.com/evenyosua18/template-project-structure/app/infrastructure/proto/pb"
	"github.com/evenyosua18/template-project-structure/app/usecase/user"
)

type ServiceUser struct {
	uc user.IUserInteraction

	pb.UnimplementedUserServiceServer
}

func NewUserService(uc user.IUserInteraction) *ServiceUser {
	return &ServiceUser{uc: uc}
}
