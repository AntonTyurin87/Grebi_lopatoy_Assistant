package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/financial_assets"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertFinancialAsset ...
func (r *repository) InsertFinancialAsset(ctx context.Context, q financial_assets.Insert) (*entity.FinancialAsset, error) {
	var res dto.FinancialAssetsSlice

	if err := Selectx(ctx, r.storage, &res, insertFinancialAssetsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertFinancialAssetsQuery(q)): %w", err)
	}

	// сохраняем только один объект за раз
	asset := res.Entity()[0]

	return asset, nil
}

func insertFinancialAssetsQuery(query financial_assets.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.FinancialAssetsTableName).Columns(
		dto.FinancialAssetsColumnMainID,
		dto.FinancialAssetsColumnFinancialAssetsType,
		dto.FinancialAssetsColumnTicker,
		dto.FinancialAssetsColumnCost,
		dto.FinancialAssetsColumnQuantity,
		dto.FinancialAssetsColumnCoupon,
		dto.FinancialAssetsColumnTime,
		dto.FinancialAssetsColumnPrice,
		dto.FinancialAssetsColumnRate,
	).
		Prefix("--InsertFinancialAssets\n")

	for _, asset := range query.FinancialAssets {
		a := dto.FinancialAssetsDtoFromEntity(asset)
		insertQuery = insertQuery.Values(
			a.GetMainID(),
			a.GetFinancialAssetsType(),
			a.GetTicker(),
			a.GetCost(),
			a.GetQuantity(),
			a.GetCoupon(),
			a.GetTime(),
			a.GetPrice(),
			a.GetRate(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.FinancialAssetsColumns, ", ")))

	return insertQuery
}
