package financial_assets

import "Grebi_lopatoy_Assistant/internal/pkg/domain/entity"

// Update ...
type Update struct {
	FinancialAssets []*entity.FinancialAsset
}

// GetFinancialAssets ...
func (s *Update) GetFinancialAssets() []*entity.FinancialAsset {
	if s == nil {
		return nil
	}

	return s.FinancialAssets
}
