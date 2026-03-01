package vkontakte

import "Grebi_lopatoy_Assistant/internal/pkg/domain/presenter"

// Handler ...
type Handler struct {
	presenter presenter.Interface
}

// NewHandler создает новый обработчик
func NewHandler(
	presenter presenter.Interface,
) *Handler {
	return &Handler{
		presenter: presenter,
	}
}
