package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/financial_assets"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateFinancialAssets ...
func (r *repository) UpdateFinancialAssets(ctx context.Context, q financial_assets.Update) ([]*entity.FinancialAsset, error) {
	var res dto.FinancialAssetsSlice

	if err := Selectx(ctx, r.storage, &res, updateFinancialAssetsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateFinancialAssetsQuery(q)): %w", err)
	}

	assets := res.Entity()

	return assets, nil
}

func updateFinancialAssetsQuery(query financial_assets.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.FinancialAssetsTableName).
		Prefix("--UpdateFinancialAssets\n")

	for _, asset := range query.FinancialAssets {
		// что дополнять
		if asset.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.FinancialAssetsColumnID: asset.ID})
		}

		// чем дополнять
		if asset.MainID != 0 {
			updateQuery = updateQuery.Set(dto.FinancialAssetsColumnMainID, asset.MainID)
		}

		if asset.FinancialAssetType != 0 {
			updateQuery = updateQuery.Set(dto.FinancialAssetsColumnFinancialAssetsType, asset.FinancialAssetType)
		}

		if asset.Ticker != "" {
			updateQuery = updateQuery.Set(dto.FinancialAssetsColumnTicker, asset.Ticker)
		}

		if asset.Cost != 0 {
			updateQuery = updateQuery.Set(dto.FinancialAssetsColumnCost, asset.Cost)
		}

		if asset.Quantity != 0 {
			updateQuery = updateQuery.Set(dto.FinancialAssetsColumnQuantity, asset.Quantity)
		}

		if asset.Coupon != 0 {
			updateQuery = updateQuery.Set(dto.FinancialAssetsColumnCoupon, asset.Coupon)
		}

		if asset.Time != 0 {
			updateQuery = updateQuery.Set(dto.FinancialAssetsColumnTime, asset.Time)
		}

		if asset.Price != 0 {
			updateQuery = updateQuery.Set(dto.FinancialAssetsColumnPrice, asset.Price)
		}

		if asset.Rate != 0 {
			updateQuery = updateQuery.Set(dto.FinancialAssetsColumnRate, asset.Rate)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.FinancialAssetsColumnID))
	}

	return updateQuery
}
