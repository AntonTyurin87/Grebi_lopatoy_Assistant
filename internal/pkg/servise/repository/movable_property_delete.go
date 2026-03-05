package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/movable_property"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteMovableProperties ...
func (r *repository) DeleteMovableProperties(ctx context.Context, q movable_property.Delete) ([]*entity.MovableProperty, error) {
	var res dto.MovableProperties

	if err := Selectx(ctx, r.storage, &res, deleteMovablePropertyQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteMovablePropertyQuery(q)): %w", err)
	}

	properties := res.Entity()

	return properties, nil
}

func deleteMovablePropertyQuery(query movable_property.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.MovablePropertyTableName).
		Prefix("--DeleteMovableProperties\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.MovablePropertyColumnID] = query.IDs
	}

	if len(query.MainIDs) > 0 {
		where[dto.MovablePropertyColumnMainID] = query.MainIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.MovablePropertyColumns))

	return deleteQuery
}
