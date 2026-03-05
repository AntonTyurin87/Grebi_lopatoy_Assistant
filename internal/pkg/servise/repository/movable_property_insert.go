package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/movable_property"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertMovableProperty ...
func (r *repository) InsertMovableProperty(ctx context.Context, q movable_property.Insert) (*entity.MovableProperty, error) {
	var res dto.MovableProperties

	if err := Selectx(ctx, r.storage, &res, insertMovablePropertyQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertMovablePropertyQuery(q)): %w", err)
	}

	// сохраняем только один объект за раз
	property := res.Entity()[0]

	return property, nil
}

func insertMovablePropertyQuery(query movable_property.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.MovablePropertyTableName).Columns(
		dto.MovablePropertyColumnMainID,
		dto.MovablePropertyColumnMovablePropertyType,
		dto.MovablePropertyColumnName,
		dto.MovablePropertyColumnCost,
		dto.MovablePropertyColumnTerm,
		dto.MovablePropertyColumnPayment,
	).
		Prefix("--InsertMovableProperty\n")

	for _, property := range query.MovableProperties {
		a := dto.MovablePropertyDtoFromEntity(property)
		insertQuery = insertQuery.Values(
			a.GetMainID(),
			a.GetMovablePropertyType(),
			a.GetName(),
			a.GetCost(),
			a.GetTerm(),
			a.GetPayment(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.MovablePropertyColumns, ", ")))

	return insertQuery
}
