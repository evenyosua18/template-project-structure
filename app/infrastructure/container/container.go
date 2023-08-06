//go:build wireinject
// +build wireinject

package container

import (
	"github.com/evenyosua18/template-project-structure/app/infrastructure/server/grpc/builder"
	userRepo "github.com/evenyosua18/template-project-structure/app/repository/example_db/user"
	"github.com/evenyosua18/template-project-structure/app/usecase/user"
	"github.com/google/wire"
)

func InitializeUserInteraction() (interaction user.IUserInteraction) {
	wire.Build(userRepo.NewUserRepository, builder.NewUserBuilder, user.NewUserInteraction)
	return
}
