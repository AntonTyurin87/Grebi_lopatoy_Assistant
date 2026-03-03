package business

import "Grebi_lopatoy_Assistant/internal/pkg/domain/entity"

// Update ...
type Update struct {
	Businesses []*entity.Business
}

// GetBusinesses ...
func (s *Update) GetBusinesses() []*entity.Business {
	if s == nil {
		return nil
	}

	return s.Businesses
}
