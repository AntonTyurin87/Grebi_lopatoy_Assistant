package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/business"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteBusinesses ...
func (r *repository) DeleteBusinesses(ctx context.Context, q business.Delete) ([]*entity.Business, error) {
	var res dto.Businesses

	if err := Selectx(ctx, r.storage, &res, deleteBusinessQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteBusinessQuery(q)): %w", err)
		//TODO подумать про логирование ошибки
	}

	business := res.Entity()

	return business, nil
}

func deleteBusinessQuery(query business.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.BusinessTableName).
		Prefix("--DeleteBusinesses\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.BusinessColumnID] = query.IDs
	}

	if len(query.MainIDs) > 0 {
		where[dto.BusinessColumnMainID] = query.MainIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.BusinessColumns))

	return deleteQuery
}
