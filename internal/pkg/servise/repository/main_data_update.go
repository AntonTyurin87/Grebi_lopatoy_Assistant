package repository

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/dto"
	"Grebi_lopatoy_Assistant/internal/pkg/servise/repository/query/main_data"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

// UpdateMainData ...
func (r *repository) UpdateMainData(ctx context.Context, q main_data.Update) ([]*entity.MainData, error) {
	var res dto.MainDataSlice

	if err := Selectx(ctx, r.storage, &res, updateMainDataQuery(q)); err != nil {
		return nil, fmt.Errorf("selectx(ctx, r.storage, &res, updateMainDataQuery(q)): %w", err)
	}

	data := res.Entity()

	return data, nil
}

func updateMainDataQuery(query main_data.Update) sq.UpdateBuilder {
	updateQuery := sq.StatementBuilder.Update(dto.MainDataTableName).
		Prefix("--UpdateMainData\n")

	for _, data := range query.MainDataSlice {
		// что дополнять
		if data.ID != 0 {
			updateQuery = updateQuery.Where(sq.Eq{dto.MainDataColumnID: data.ID})
		}

		// чем дополнять
		if data.ChatID != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnChatID, data.ChatID)
		}

		if data.UserID != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnUserID, data.UserID)
		}

		if data.Step != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnStep, data.Step)
		}

		updateQuery = updateQuery.Set(dto.MainDataColumnIsCreated, data.IsCreated)

		if data.Profession != "" {
			updateQuery = updateQuery.Set(dto.MainDataColumnProfession, data.Profession)
		}

		if data.Gender != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnGender, data.Gender)
		}

		if data.World != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnWorld, data.World)
		}

		updateQuery = updateQuery.Set(dto.MainDataColumnMarriage, data.Marriage)

		if data.Children != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnChildren, data.Children)
		}

		if data.Wishes != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnWishes, data.Wishes)
		}

		if data.Turn != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnTurn, data.Turn)
		}

		if data.Salary != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnSalary, data.Salary)
		}

		if data.SalaryExtraName != "" {
			updateQuery = updateQuery.Set(dto.MainDataColumnSalaryExtraName, data.SalaryExtraName)
		}

		if data.SalaryExtra != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnSalaryExtra, data.SalaryExtra)
		}

		if data.CostHouse != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnCostHouse, data.CostHouse)
		}

		if data.CostFood != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnCostFood, data.CostFood)
		}

		if data.CostTransport != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnCostTransport, data.CostTransport)
		}

		if data.CostCloth != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnCostCloth, data.CostCloth)
		}

		if data.CostExtraName != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnCostExtraName, data.CostExtraName)
		}

		if data.CostExtra != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnCostExtra, data.CostExtra)
		}

		if data.TotalIncome != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnTotalIncome, data.TotalIncome)
		}

		if data.TotalOutcome != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnTotalOutcome, data.TotalOutcome)
		}

		if data.Flow != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnFlow, data.Flow)
		}

		if data.Cash != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnCash, data.Cash)
		}

		if data.MenuID != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnMenuID, data.MenuID)
		}

		if data.LastActiveID != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnLastActiveID, data.LastActiveID)
		}

		if data.Debt != 0 {
			updateQuery = updateQuery.Set(dto.MainDataColumnDebt, data.Debt)
		}

		if data.CreatedAt != "" {
			updateQuery = updateQuery.Set(dto.MainDataColumnCreatedAt, data.CreatedAt)
		}

		if data.UpdatedAt != "" {
			updateQuery = updateQuery.Set(dto.MainDataColumnUpdatedAt, data.UpdatedAt)
		}

		updateQuery = updateQuery.Suffix(fmt.Sprintf("RETURNING %s", dto.MainDataColumnID))
	}

	return updateQuery
}
