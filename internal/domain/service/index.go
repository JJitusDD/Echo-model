package service

import "echo-model/internal/infrastructure/facade"

type Service struct {
	*facade.EchoModelFacade
}

func NewService(f *facade.EchoModelFacade) *Service {
	return &Service{
		f,
	}
}
