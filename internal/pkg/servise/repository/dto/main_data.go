package dto

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
)

const (
	MainDataTableName = "main_data"

	MainDataColumnID              = "id"
	MainDataColumnChatID          = "chat_id"
	MainDataColumnUserID          = "user_id"
	MainDataColumnStep            = "step"
	MainDataColumnIsCreated       = "is_created"
	MainDataColumnProfession      = "profession"
	MainDataColumnGender          = "gender"
	MainDataColumnWorld           = "world"
	MainDataColumnMarriage        = "marriage"
	MainDataColumnChildren        = "children"
	MainDataColumnWishes          = "wishes"
	MainDataColumnTurn            = "turn"
	MainDataColumnSalary          = "salary"
	MainDataColumnSalaryExtraName = "salary_extra_name"
	MainDataColumnSalaryExtra     = "salary_extra"
	MainDataColumnCostHouse       = "cost_house"
	MainDataColumnCostFood        = "cost_food"
	MainDataColumnCostTransport   = "cost_transport"
	MainDataColumnCostCloth       = "cost_cloth"
	MainDataColumnCostExtraName   = "cost_extra_name"
	MainDataColumnCostExtra       = "cost_extra"
	MainDataColumnTotalIncome     = "total_income"
	MainDataColumnTotalOutcome    = "total_outcome"
	MainDataColumnFlow            = "flow"
	MainDataColumnCash            = "cash"
	MainDataColumnMenuID          = "menu_id"
	MainDataColumnLastActiveID    = "last_active_id"
	MainDataColumnDebt            = "debt"
	MainDataColumnCreatedAt       = "created_at"
	MainDataColumnUpdatedAt       = "updated_at"
)

var MainDataColumns = []string{
	MainDataColumnID,
	MainDataColumnChatID,
	MainDataColumnUserID,
	MainDataColumnStep,
	MainDataColumnIsCreated,
	MainDataColumnProfession,
	MainDataColumnGender,
	MainDataColumnWorld,
	MainDataColumnMarriage,
	MainDataColumnChildren,
	MainDataColumnWishes,
	MainDataColumnTurn,
	MainDataColumnSalary,
	MainDataColumnSalaryExtraName,
	MainDataColumnSalaryExtra,
	MainDataColumnCostHouse,
	MainDataColumnCostFood,
	MainDataColumnCostTransport,
	MainDataColumnCostCloth,
	MainDataColumnCostExtraName,
	MainDataColumnCostExtra,
	MainDataColumnTotalIncome,
	MainDataColumnTotalOutcome,
	MainDataColumnFlow,
	MainDataColumnCash,
	MainDataColumnMenuID,
	MainDataColumnLastActiveID,
	MainDataColumnDebt,
	MainDataColumnCreatedAt,
	MainDataColumnUpdatedAt,
}

// MainData ...
type MainData struct {
	ID              int64  `json:"id"`
	ChatID          int64  `json:"chatID"`
	UserID          int64  `json:"userID"`
	Step            int64  `json:"step"`
	IsCreated       bool   `json:"isCreated"`
	Profession      string `json:"profession"`
	Gender          int64  `json:"gender"`
	World           int64  `json:"world"`
	Marriage        bool   `json:"marriage"`
	Children        int64  `json:"children"`
	Wishes          int64  `json:"wishes"`
	Turn            int64  `json:"turn"`
	Salary          int64  `json:"salary"`
	SalaryExtraName string `json:"salaryExtraName"`
	SalaryExtra     int64  `json:"salaryExtra"`
	CostHouse       int64  `json:"costHouse"`
	CostFood        int64  `json:"costFood"`
	CostTransport   int64  `json:"costTransport"`
	CostCloth       int64  `json:"costCloth"`
	CostExtraName   int64  `json:"costExtraName"`
	CostExtra       int64  `json:"costExtra"`
	TotalIncome     int64  `json:"totalIncome"`
	TotalOutcome    int64  `json:"totalOutcome"`
	Flow            int64  `json:"flow"`
	Cash            int64  `json:"cash"`
	MenuID          int64  `json:"menuID"`
	LastActiveID    int64  `json:"lastActiveID"`
	Debt            int64  `json:"debt"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

// GetID ...
func (m *MainData) GetID() int64 {
	if m == nil {
		return 0
	}
	return m.ID
}

// GetChatID ...
func (m *MainData) GetChatID() int64 {
	if m == nil {
		return 0
	}
	return m.ChatID
}

// GetUserID ...
func (m *MainData) GetUserID() int64 {
	if m == nil {
		return 0
	}
	return m.UserID
}

// GetStep ...
func (m *MainData) GetStep() int64 {
	if m == nil {
		return 0
	}
	return m.Step
}

// GetIsCreated ...
func (m *MainData) GetIsCreated() bool {
	if m == nil {
		return false
	}
	return m.IsCreated
}

// GetProfession ...
func (m *MainData) GetProfession() string {
	if m == nil {
		return ""
	}
	return m.Profession
}

// GetGender ...
func (m *MainData) GetGender() int64 {
	if m == nil {
		return 0
	}
	return m.Gender
}

// GetWorld ...
func (m *MainData) GetWorld() int64 {
	if m == nil {
		return 0
	}
	return m.World
}

// GetMarriage ...
func (m *MainData) GetMarriage() bool {
	if m == nil {
		return false
	}
	return m.Marriage
}

// GetChildren ...
func (m *MainData) GetChildren() int64 {
	if m == nil {
		return 0
	}
	return m.Children
}

// GetWishes ...
func (m *MainData) GetWishes() int64 {
	if m == nil {
		return 0
	}
	return m.Wishes
}

// GetTurn ...
func (m *MainData) GetTurn() int64 {
	if m == nil {
		return 0
	}
	return m.Turn
}

// GetSalary ...
func (m *MainData) GetSalary() int64 {
	if m == nil {
		return 0
	}
	return m.Salary
}

// GetSalaryExtraName ...
func (m *MainData) GetSalaryExtraName() string {
	if m == nil {
		return ""
	}
	return m.SalaryExtraName
}

// GetSalaryExtra ...
func (m *MainData) GetSalaryExtra() int64 {
	if m == nil {
		return 0
	}
	return m.SalaryExtra
}

// GetCostHouse ...
func (m *MainData) GetCostHouse() int64 {
	if m == nil {
		return 0
	}
	return m.CostHouse
}

// GetCostFood ...
func (m *MainData) GetCostFood() int64 {
	if m == nil {
		return 0
	}
	return m.CostFood
}

// GetCostTransport ...
func (m *MainData) GetCostTransport() int64 {
	if m == nil {
		return 0
	}
	return m.CostTransport
}

// GetCostCloth ...
func (m *MainData) GetCostCloth() int64 {
	if m == nil {
		return 0
	}
	return m.CostCloth
}

// GetCostExtraName ...
func (m *MainData) GetCostExtraName() int64 {
	if m == nil {
		return 0
	}
	return m.CostExtraName
}

// GetCostExtra ...
func (m *MainData) GetCostExtra() int64 {
	if m == nil {
		return 0
	}
	return m.CostExtra
}

// GetTotalIncome ...
func (m *MainData) GetTotalIncome() int64 {
	if m == nil {
		return 0
	}
	return m.TotalIncome
}

// GetTotalOutcome ...
func (m *MainData) GetTotalOutcome() int64 {
	if m == nil {
		return 0
	}
	return m.TotalOutcome
}

// GetFlow ...
func (m *MainData) GetFlow() int64 {
	if m == nil {
		return 0
	}
	return m.Flow
}

// GetCash ...
func (m *MainData) GetCash() int64 {
	if m == nil {
		return 0
	}
	return m.Cash
}

// GetMenuID ...
func (m *MainData) GetMenuID() int64 {
	if m == nil {
		return 0
	}
	return m.MenuID
}

// GetLastActiveID ...
func (m *MainData) GetLastActiveID() int64 {
	if m == nil {
		return 0
	}
	return m.LastActiveID
}

// GetDebt ...
func (m *MainData) GetDebt() int64 {
	if m == nil {
		return 0
	}
	return m.Debt
}

// GetCreatedAt ...
func (m *MainData) GetCreatedAt() string {
	if m == nil {
		return ""
	}
	return m.CreatedAt
}

// GetUpdatedAt ...
func (m *MainData) GetUpdatedAt() string {
	if m == nil {
		return ""
	}
	return m.UpdatedAt
}

// Entity ...
func (m *MainData) Entity() *entity.MainData {
	if m == nil {
		return nil
	}

	return &entity.MainData{
		ID:              m.GetID(),
		ChatID:          m.GetChatID(),
		UserID:          m.GetUserID(),
		Step:            entity.Step(m.GetStep()),
		IsCreated:       m.GetIsCreated(),
		Profession:      m.GetProfession(),
		Gender:          entity.Gender(m.GetGender()),
		World:           entity.World(m.GetWorld()),
		Marriage:        m.GetMarriage(),
		Children:        m.GetChildren(),
		Wishes:          m.GetWishes(),
		Turn:            m.GetTurn(),
		Salary:          m.GetSalary(),
		SalaryExtraName: m.GetSalaryExtraName(),
		SalaryExtra:     m.GetSalaryExtra(),
		CostHouse:       m.GetCostHouse(),
		CostFood:        m.GetCostFood(),
		CostTransport:   m.GetCostTransport(),
		CostCloth:       m.GetCostCloth(),
		CostExtraName:   m.GetCostExtraName(),
		CostExtra:       m.GetCostExtra(),
		TotalIncome:     m.GetTotalIncome(),
		TotalOutcome:    m.GetTotalOutcome(),
		Flow:            m.GetFlow(),
		Cash:            m.GetCash(),
		MenuID:          m.GetMenuID(),
		LastActiveID:    m.GetLastActiveID(),
		Debt:            m.GetDebt(),
		CreatedAt:       m.GetCreatedAt(),
		UpdatedAt:       m.GetUpdatedAt(),
	}
}

// MainDataSlice ...
type MainDataSlice []*MainData

// Entity ...
func (m MainDataSlice) Entity() []*entity.MainData {
	result := make([]*entity.MainData, 0, len(m))
	for _, item := range m {
		if item != nil {
			result = append(result, item.Entity())
		}
	}
	return result
}

// MainDataDtoFromEntity ...
func MainDataDtoFromEntity(e *entity.MainData) *MainData {
	if e == nil {
		return nil
	}
	return &MainData{
		ID:              e.GetID(),
		ChatID:          e.GetChatID(),
		UserID:          e.GetUserID(),
		Step:            int64(e.GetStep()),
		IsCreated:       e.GetIsCreated(),
		Profession:      e.GetProfession(),
		Gender:          int64(e.GetGender()),
		World:           int64(e.GetWorld()),
		Marriage:        e.GetMarriage(),
		Children:        e.GetChildren(),
		Wishes:          e.GetWishes(),
		Turn:            e.GetTurn(),
		Salary:          e.GetSalary(),
		SalaryExtraName: e.GetSalaryExtraName(),
		SalaryExtra:     e.GetSalaryExtra(),
		CostHouse:       e.GetCostHouse(),
		CostFood:        e.GetCostFood(),
		CostTransport:   e.GetCostTransport(),
		CostCloth:       e.GetCostCloth(),
		CostExtraName:   e.GetCostExtraName(),
		CostExtra:       e.GetCostExtra(),
		TotalIncome:     e.GetTotalIncome(),
		TotalOutcome:    e.GetTotalOutcome(),
		Flow:            e.GetFlow(),
		Cash:            e.GetCash(),
		MenuID:          e.GetMenuID(),
		LastActiveID:    e.GetLastActiveID(),
		Debt:            e.GetDebt(),
		CreatedAt:       e.GetCreatedAt(),
		UpdatedAt:       e.GetUpdatedAt(),
	}
}

// MainDataDtoFromEntitySlice ...
func MainDataDtoFromEntitySlice(entities []*entity.MainData) MainDataSlice {
	if entities == nil {
		return nil
	}
	result := make(MainDataSlice, 0, len(entities))
	for _, e := range entities {
		result = append(result, MainDataDtoFromEntity(e))
	}
	return result
}
