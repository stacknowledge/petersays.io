package prophet

import "github.com/stacknowledge/petersays.io/component/domain"

type (
	Service interface {
		Prophesize() string
		Enlightment(string) (string, error)
	}

	ProphetService struct {
	}
)

func (example ProphetService) Prophesize() string {

	domainObject := new(domain.Saying)
	domainObject.Text = "A petersay saying!"

	return domainObject.Text
}

func (example ProphetService) Enlightment(enlightment string) (string, error) {

	domainObject := new(domain.Saying)
	domainObject.Text = enlightment

	return domainObject.Text, nil
}
