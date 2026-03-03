package movable_property

import "Grebi_lopatoy_Assistant/internal/pkg/domain/entity"

// Update ...
type Update struct {
	MovableProperties []*entity.MovableProperty
}

// GetMovableProperties ...
func (s *Update) GetMovableProperties() []*entity.MovableProperty {
	if s == nil {
		return nil
	}

	return s.MovableProperties
}
