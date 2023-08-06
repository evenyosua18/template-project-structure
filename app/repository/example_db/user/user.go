package user

type IUserRepository interface {
	GetUser()
}

type RepositoryUser struct {
}

func NewUserRepository() IUserRepository {
	return &RepositoryUser{}
}
