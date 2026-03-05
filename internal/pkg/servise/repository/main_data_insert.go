package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/main_data"
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// InsertMainData ...
func (r *repository) InsertMainData(ctx context.Context, q main_data.Insert) (*entity.MainData, error) {
	var res dto.MainDataSlice

	if err := Selectx(ctx, r.storage, &res, insertMainDataQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, insertMainDataQuery(q)): %w", err)
	}

	// сохраняем только один объект за раз
	data := res.Entity()[0]

	return data, nil
}

func insertMainDataQuery(query main_data.Insert) sq.InsertBuilder {
	insertQuery := sq.StatementBuilder.Insert(dto.MainDataTableName).Columns(
		dto.MainDataColumnChatID,
		dto.MainDataColumnUserID,
		dto.MainDataColumnStep,
		dto.MainDataColumnIsCreated,
		dto.MainDataColumnProfession,
		dto.MainDataColumnGender,
		dto.MainDataColumnWorld,
		dto.MainDataColumnMarriage,
		dto.MainDataColumnChildren,
		dto.MainDataColumnWishes,
		dto.MainDataColumnTurn,
		dto.MainDataColumnSalary,
		dto.MainDataColumnSalaryExtraName,
		dto.MainDataColumnSalaryExtra,
		dto.MainDataColumnCostHouse,
		dto.MainDataColumnCostFood,
		dto.MainDataColumnCostTransport,
		dto.MainDataColumnCostCloth,
		dto.MainDataColumnCostExtraName,
		dto.MainDataColumnCostExtra,
		dto.MainDataColumnTotalIncome,
		dto.MainDataColumnTotalOutcome,
		dto.MainDataColumnFlow,
		dto.MainDataColumnCash,
		dto.MainDataColumnMenuID,
		dto.MainDataColumnLastActiveID,
		dto.MainDataColumnDebt,
		dto.MainDataColumnCreatedAt,
		dto.MainDataColumnUpdatedAt,
	).
		Prefix("--InsertMainData\n")

	for _, data := range query.MainDataSlice {
		a := dto.MainDataDtoFromEntity(data)
		insertQuery = insertQuery.Values(
			a.GetChatID(),
			a.GetUserID(),
			a.GetStep(),
			a.GetIsCreated(),
			a.GetProfession(),
			a.GetGender(),
			a.GetWorld(),
			a.GetMarriage(),
			a.GetChildren(),
			a.GetWishes(),
			a.GetTurn(),
			a.GetSalary(),
			a.GetSalaryExtraName(),
			a.GetSalaryExtra(),
			a.GetCostHouse(),
			a.GetCostFood(),
			a.GetCostTransport(),
			a.GetCostCloth(),
			a.GetCostExtraName(),
			a.GetCostExtra(),
			a.GetTotalIncome(),
			a.GetTotalOutcome(),
			a.GetFlow(),
			a.GetCash(),
			a.GetMenuID(),
			a.GetLastActiveID(),
			a.GetDebt(),
			a.GetCreatedAt(),
			a.GetUpdatedAt(),
		)
	}

	insertQuery = insertQuery.Suffix(fmt.Sprintf("RETURNING %s", strings.Join(dto.MainDataColumns, ", ")))

	return insertQuery
}
