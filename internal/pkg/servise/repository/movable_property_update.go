package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/movable_property"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateMovableProperties ...
func (r *repository) UpdateMovableProperties(ctx context.Context, q movable_property.Update) ([]*entity.MovableProperty, error) {
	var res dto.MovableProperties

	if err := Selectx(ctx, r.storage, &res, updateMovablePropertyQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateMovablePropertyQuery(q)): %w", err)
	}

	properties := res.Entity()

	return properties, nil
}

func updateMovablePropertyQuery(query movable_property.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.MovablePropertyTableName).
		Prefix("--UpdateMovableProperties\n")

	for _, property := range query.MovableProperties {
		// что дополнять
		if property.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.MovablePropertyColumnID: property.ID})
		}

		// чем дополнять
		if property.MainID != 0 {
			updateQuery = updateQuery.Set(dto.MovablePropertyColumnMainID, property.MainID)
		}

		if property.MovablePropertyType != 0 {
			updateQuery = updateQuery.Set(dto.MovablePropertyColumnMovablePropertyType, property.MovablePropertyType)
		}

		if property.Name != "" {
			updateQuery = updateQuery.Set(dto.MovablePropertyColumnName, property.Name)
		}

		if property.Cost != 0 {
			updateQuery = updateQuery.Set(dto.MovablePropertyColumnCost, property.Cost)
		}

		if property.Term != 0 {
			updateQuery = updateQuery.Set(dto.MovablePropertyColumnTerm, property.Term)
		}

		if property.Payment != 0 {
			updateQuery = updateQuery.Set(dto.MovablePropertyColumnPayment, property.Payment)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.MovablePropertyColumnID))
	}

	return updateQuery
}
