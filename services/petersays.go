package services

type (
	PetersaysInterface interface {
		PeterSay() string
	}

	PetersaysService struct {
	}
)

func (service *PetersaysService) PeterSay() string {

	return ""
}
