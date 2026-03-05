package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/movable_property"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectMovableProperties ...
func (r *repository) SelectMovableProperties(ctx context.Context, q movable_property.Select) ([]*entity.MovableProperty, error) {
	var res dto.MovableProperties

	if err := Selectx(ctx, r.storage, &res, selectMovablePropertyQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectMovablePropertyQuery(q)): %w", err)
	}

	properties := res.Entity()

	return properties, nil
}

func selectMovablePropertyQuery(query movable_property.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.MovablePropertyColumns...).
		From(dto.MovablePropertyTableName).
		Prefix("--SelectMovableProperties\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.MovablePropertyColumnID] = query.IDs
	}

	if len(query.MainIDs) > 0 {
		where[dto.MovablePropertyColumnMainID] = query.MainIDs
	}

	// ищем только по ID и MainID
	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.MovablePropertyColumnID)

	return selectQuery
}
