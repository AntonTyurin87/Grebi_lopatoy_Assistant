package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/real_estate"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectRealEstates ...
func (r *repository) SelectRealEstates(ctx context.Context, q real_estate.Select) ([]*entity.RealEstate, error) {
	var res dto.RealEstates

	if err := Selectx(ctx, r.storage, &res, selectRealEstateQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectRealEstateQuery(q)): %w", err)
	}

	properties := res.Entity()

	return properties, nil
}

func selectRealEstateQuery(query real_estate.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.RealEstateColumns...).
		From(dto.RealEstateTableName).
		Prefix("--SelectRealEstates\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.RealEstateColumnID] = query.IDs
	}

	if len(query.MainIDs) > 0 {
		where[dto.RealEstateColumnMainID] = query.MainIDs
	}

	// ищем только по ID и MainID
	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.RealEstateColumnID)

	return selectQuery
}
