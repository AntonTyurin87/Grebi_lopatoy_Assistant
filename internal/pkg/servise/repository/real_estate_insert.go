package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/real_estate"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertRealEstate ...
func (r *repository) InsertRealEstate(ctx context.Context, q real_estate.Insert) (*entity.RealEstate, error) {
	var res dto.RealEstates

	if err := Selectx(ctx, r.storage, &res, insertRealEstateQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertRealEstateQuery(q)): %w", err)
	}

	// сохраняем только один объект за раз
	property := res.Entity()[0]

	return property, nil
}

func insertRealEstateQuery(query real_estate.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.RealEstateTableName).Columns(
		dto.RealEstateColumnMainID,
		dto.RealEstateColumnRealEstateType,
		dto.RealEstateColumnName,
		dto.RealEstateColumnCost,
		dto.RealEstateColumnTerm,
		dto.RealEstateColumnPayment,
	).
		Prefix("--InsertRealEstate\n")

	for _, property := range query.RealEstates {
		a := dto.RealEstateDtoFromEntity(property)
		insertQuery = insertQuery.Values(
			a.GetMainID(),
			a.GetRealEstateType(),
			a.GetName(),
			a.GetCost(),
			a.GetTerm(),
			a.GetPayment(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.RealEstateColumns, ", ")))

	return insertQuery
}
