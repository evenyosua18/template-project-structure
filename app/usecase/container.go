package usecase

import (
	"github.com/evenyosua18/template-project-structure/app/infrastructure/server/grpc/builder"
	"github.com/evenyosua18/template-project-structure/app/repository/example_db/example"
	example2 "github.com/evenyosua18/template-project-structure/app/usecase/example"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, err := di.NewBuilder()

	if err != nil {
		panic(err)
	}

	if err = builder.Add([]di.Def{
		//input all use case function here
		//ex : {Name: "", Build: func},
		{Name: "exampleCTN", Build: exampleInteraction},
	}...); err != nil {
		panic(err)
	}

	return &Container{ctn: builder.Build()}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func exampleInteraction(_ di.Container) (interface{}, error) {
	repo := example.NewExampleRepository()
	out := &builder.ExampleBuilder{}
	return example2.NewInteractionExample(repo, out), nil
}
