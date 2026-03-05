package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/real_estate"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateRealEstates ...
func (r *repository) UpdateRealEstates(ctx context.Context, q real_estate.Update) ([]*entity.RealEstate, error) {
	var res dto.RealEstates

	if err := Selectx(ctx, r.storage, &res, updateRealEstateQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateRealEstateQuery(q)): %w", err)
	}

	properties := res.Entity()

	return properties, nil
}

func updateRealEstateQuery(query real_estate.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.RealEstateTableName).
		Prefix("--UpdateRealEstates\n")

	for _, property := range query.RealEstates {
		// что дополнять
		if property.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.RealEstateColumnID: property.ID})
		}

		// чем дополнять
		if property.MainID != 0 {
			updateQuery = updateQuery.Set(dto.RealEstateColumnMainID, property.MainID)
		}

		if property.RealEstateType != 0 {
			updateQuery = updateQuery.Set(dto.RealEstateColumnRealEstateType, property.RealEstateType)
		}

		if property.Name != "" {
			updateQuery = updateQuery.Set(dto.RealEstateColumnName, property.Name)
		}

		if property.Cost != 0 {
			updateQuery = updateQuery.Set(dto.RealEstateColumnCost, property.Cost)
		}

		if property.Term != 0 {
			updateQuery = updateQuery.Set(dto.RealEstateColumnTerm, property.Term)
		}

		if property.Payment != 0 {
			updateQuery = updateQuery.Set(dto.RealEstateColumnPayment, property.Payment)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.RealEstateColumnID))
	}

	return updateQuery
}
