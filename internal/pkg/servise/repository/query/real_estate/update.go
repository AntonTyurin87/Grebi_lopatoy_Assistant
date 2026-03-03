package real_estate

import "Grebi_lopatoy_Assistant/internal/pkg/domain/entity"

// Update ...
type Update struct {
	RealEstates []*entity.RealEstate
}

// GetRealEstates ...
func (s *Update) GetRealEstates() []*entity.RealEstate {
	if s == nil {
		return nil
	}

	return s.RealEstates
}
