package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/business"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateBusiness ...
func (r *repository) UpdateBusiness(ctx context.Context, q business.Update) ([]*entity.Business, error) {
	var res dto.Businesses

	if err := Selectx(ctx, r.storage, &res, updateBusinessQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateBusinessQuery(q)): %w", err)
	}

	business := res.Entity()

	return business, nil
}

func updateBusinessQuery(query business.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.BusinessTableName).
		Prefix("--UpdateBusiness\n")

	for _, business := range query.Businesses {

		// что дополнять
		if business.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.BusinessColumnID: business.ID})
		}

		// чем дополнять
		if business.MainID != 0 {
			updateQuery = updateQuery.Set(dto.BusinessColumnMainID, business.MainID)
		}

		if business.BusinessType != 0 {
			updateQuery = updateQuery.Set(dto.BusinessColumnBusinessType, business.BusinessType)
		}

		if business.Name != "" {
			updateQuery = updateQuery.Set(dto.BusinessColumnName, business.Name)
		}

		if business.Cost != 0 {
			updateQuery = updateQuery.Set(dto.BusinessColumnCost, business.Cost)
		}

		if business.Profit != 0 {
			updateQuery = updateQuery.Set(dto.BusinessColumnProfit, business.Profit)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.BusinessColumnID))
	}

	return updateQuery
}
