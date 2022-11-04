package example

import "github.com/evenyosua18/template-project-structure/app/usecase/example"

type ServiceExample struct {
	uc example.InputPortExample
}

func NewServiceExample(uc example.InputPortExample) *ServiceExample {
	return &ServiceExample{uc: uc}
}
