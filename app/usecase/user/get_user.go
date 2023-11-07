package user

import "context"

func (i *InteractionUser) GetUser(ctx context.Context, in interface{}) (interface{}, error) {
	return i.out.GetUserResponse(nil)
}
