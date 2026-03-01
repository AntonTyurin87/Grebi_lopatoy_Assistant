package usecase

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/presenter"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository"
)

type usecase struct {
	presenter presenter.Interface

	repository repository.Repository
}

// NewUsacase ...
func NewUsacase(
	presenter presenter.Interface,

	repository repository.Repository,
) *usecase {
	return &usecase{
		presenter:  presenter,
		repository: repository,
	}
}
