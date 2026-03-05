package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/real_estate"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteRealEstates ...
func (r *repository) DeleteRealEstates(ctx context.Context, q real_estate.Delete) ([]*entity.RealEstate, error) {
	var res dto.RealEstates

	if err := Selectx(ctx, r.storage, &res, deleteRealEstateQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteRealEstateQuery(q)): %w", err)
	}

	properties := res.Entity()

	return properties, nil
}

func deleteRealEstateQuery(query real_estate.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.RealEstateTableName).
		Prefix("--DeleteRealEstates\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.RealEstateColumnID] = query.IDs
	}

	if len(query.MainIDs) > 0 {
		where[dto.RealEstateColumnMainID] = query.MainIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.RealEstateColumns))

	return deleteQuery
}
