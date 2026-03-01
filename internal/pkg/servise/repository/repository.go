package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/servise/storage"
)

type repository struct {
	storage storage.Storage
}

func NewRepository(
	storage storage.Storage,
) Repository {
	return &repository{
		storage: storage,
	}
}
