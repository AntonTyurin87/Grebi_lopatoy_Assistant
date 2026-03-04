package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/business"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectBusiness ...
func (r *repository) SelectBusiness(ctx context.Context, q business.Select) ([]*entity.Business, error) {
	var res dto.Businesses

	if err := Selectx(ctx, r.storage, &res, selectBusinessQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectBusinessQuery(q)): %w", err)
	}

	business := res.Entity()

	return business, nil
}

func selectBusinessQuery(query business.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.BusinessColumns...).
		From(dto.BusinessTableName).
		Prefix("--SelectBusiness\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.BusinessColumnID] = query.IDs
	}

	if len(query.MainIDs) > 0 {
		where[dto.BusinessColumnMainID] = query.MainIDs
	}

	//ищем только по ID и MainID
	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.BusinessColumnID)

	return selectQuery
}
