package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/main_data"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// SelectMainData ...
func (r *repository) SelectMainData(ctx context.Context, q main_data.Select) ([]*entity.MainData, error) {
	var res dto.MainDataSlice

	if err := Selectx(ctx, r.storage, &res, selectMainDataQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, selectMainDataQuery(q)): %w", err)
	}

	data := res.Entity()

	return data, nil
}

func selectMainDataQuery(query main_data.Select) sq.SelectBuilder {
	selectQuery := sq.StatementBuilder.Select(dto.MainDataColumns...).
		From(dto.MainDataTableName).
		Prefix("--SelectMainData\n")

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

	selectQuery = selectQuery.Where(where)
	selectQuery = selectQuery.OrderBy(dto.MainDataColumnID)

	return selectQuery
}
