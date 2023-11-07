package user

import "context"

type IUserRepository interface {
	GetUser(ctx context.Context, in interface{}) (interface{}, error)
}

type RepositoryUser struct {
}

func NewUserRepository() IUserRepository {
	return &RepositoryUser{}
}
