package example

import "github.com/stacknowledge/petersays.io/component/domain"

type (
	Service interface {
		Method() string
	}

	ExampleService struct {
	}
)

func (example ExampleService) Method() string {

	domainObject := new(domain.Example)
	domainObject.Text = "Method string Example!"

	return domainObject.Text
}
