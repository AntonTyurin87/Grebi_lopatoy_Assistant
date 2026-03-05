package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/main_data"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// DeleteMainData ...
func (r *repository) DeleteMainData(ctx context.Context, q main_data.Delete) ([]*entity.MainData, error) {
	var res dto.MainDataSlice

	if err := Selectx(ctx, r.storage, &res, deleteMainDataQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, deleteMainDataQuery(q)): %w", err)
	}

	data := res.Entity()

	return data, nil
}

func deleteMainDataQuery(query main_data.Delete) sq.DeleteBuilder {
	deleteQuery := sq.StatementBuilder.Delete(dto.MainDataTableName).
		Prefix("--DeleteMainData\n")

	where := sq.Eq{}

	if len(query.IDs) > 0 {
		where[dto.MainDataColumnID] = query.IDs
	}

	if len(query.ChatIDs) > 0 {
		where[dto.MainDataColumnChatID] = query.ChatIDs
	}

	if len(query.UserIDs) > 0 {
		where[dto.MainDataColumnUserID] = query.UserIDs
	}

	deleteQuery = deleteQuery.Where(where)

	deleteQuery = deleteQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.MainDataColumns))

	return deleteQuery
}
