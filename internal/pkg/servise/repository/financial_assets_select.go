package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/financial_assets"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectFinancialAssets ...
func (r *repository) SelectFinancialAssets(ctx context.Context, q financial_assets.Select) ([]*entity.FinancialAsset, error) {
	var res dto.FinancialAssetsSlice

	if err := Selectx(ctx, r.storage, &res, selectFinancialAssetsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectFinancialAssetsQuery(q)): %w", err)
	}

	assets := res.Entity()

	return assets, nil
}

func selectFinancialAssetsQuery(query financial_assets.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.FinancialAssetsColumns...).
		From(dto.FinancialAssetsTableName).
		Prefix("--SelectFinancialAssets\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.FinancialAssetsColumnID] = query.IDs
	}

	if len(query.MainIDs) > 0 {
		where[dto.FinancialAssetsColumnMainID] = query.MainIDs
	}

	// ищем только по ID и MainID
	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.FinancialAssetsColumnID)

	return selectQuery
}
