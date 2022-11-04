package repository

import "context"

type ExampleRepository interface {
	GetExample(ctx context.Context, in interface{}) error
}
