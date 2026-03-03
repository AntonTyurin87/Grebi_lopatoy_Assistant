package main_data

import "Grebi_lopatoy_Assistant/internal/pkg/domain/entity"

// Update ...
type Update struct {
	MainDataSlice []*entity.MainData
}

// GetMainDataSlice ...
func (s *Update) GetMainDataSlice() []*entity.MainData {
	if s == nil {
		return nil
	}

	return s.MainDataSlice
}
