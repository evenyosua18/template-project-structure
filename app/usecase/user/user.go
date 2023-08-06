package user

import (
	"context"
	"github.com/evenyosua18/template-project-structure/app/repository/example_db/user"
)

type IUserInteraction interface {
	GetUser(ctx context.Context, in interface{}) (interface{}, error)
}

type IUserInteractionResponse interface {
	GetUserResponse(in interface{}) (interface{}, error)
}

type InteractionUser struct {
	repo user.IUserRepository
	out  IUserInteractionResponse
}

func NewUserInteraction(repo user.IUserRepository, out IUserInteractionResponse) IUserInteraction {
	return &InteractionUser{repo: repo, out: out}
}
