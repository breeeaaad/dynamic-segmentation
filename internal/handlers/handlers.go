package handlers

import "github.com/breeeaaad/dynamic-segmentation/internal/repository"

type Handlers struct {
	s *repository.Service
}

func New(s *repository.Service) *Handlers {
	return &Handlers{
		s: s,
	}
}
