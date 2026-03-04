package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/business"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertBusiness ...
func (r *repository) InsertBusiness(ctx context.Context, q business.Insert) (*entity.Business, error) {
	var res dto.Businesses

	if err := Selectx(ctx, r.storage, &res, insertBusinessQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertBusinessQuery(q)): %w", err)
		//TODO подумать про логирование ошибки
	}

	// сохраняем только один объект за раз
	business := res.Entity()[0]

	return business, nil
}

func insertBusinessQuery(query business.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.BusinessTableName).Columns(
		dto.BusinessColumnMainID,
		dto.BusinessColumnBusinessType,
		dto.BusinessColumnName,
		dto.BusinessColumnCost,
		dto.BusinessColumnProfit,
	).
		Prefix("--InsertBusiness\n")

	for _, business := range query.Businesses {
		a := dto.BusinessDtoFromEntity(business)
		insertQuery = insertQuery.Values(
			a.GetMainID(),
			a.GetBusinessType(),
			a.GetName(),
			a.GetCost(),
			a.GetProfit(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.BusinessColumns, ", ")))

	return insertQuery
}
