package example

import "context"

func (r *RepositoryExample) GetExample(ctx context.Context, in interface{}) error {
	//retrieve data from db, redis, edc here
	return nil
}
