package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/financial_assets"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteFinancialAssets ...
func (r *repository) DeleteFinancialAssets(ctx context.Context, q financial_assets.Delete) ([]*entity.FinancialAsset, error) {
	var res dto.FinancialAssetsSlice

	if err := Selectx(ctx, r.storage, &res, deleteFinancialAssetsQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteFinancialAssetsQuery(q)): %w", err)
	}

	assets := res.Entity()

	return assets, nil
}

func deleteFinancialAssetsQuery(query financial_assets.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.FinancialAssetsTableName).
		Prefix("--DeleteFinancialAssets\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.FinancialAssetsColumnID] = query.IDs
	}

	if len(query.MainIDs) > 0 {
		where[dto.FinancialAssetsColumnMainID] = query.MainIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.FinancialAssetsColumns))

	return deleteQuery
}
