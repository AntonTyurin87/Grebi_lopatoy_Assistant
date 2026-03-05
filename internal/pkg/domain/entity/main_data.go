package entity

// MainData ...
type MainData struct {
	ID              int64  `json:"id"`
	ChatID          int64  `json:"chatID"`
	UserID          int64  `json:"userID"`
	Step            Step   `json:"step"`
	IsCreated       bool   `json:"isCreated"`
	Profession      string `json:"profession"`
	Gender          Gender `json:"gender"`
	World           World  `json:"world"`
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
	MenuID          int64  `json:"menuID"`       // TODO зачем это нужно? аналог Step
	LastActiveID    int64  `json:"lastActiveID"` // TODO зачем это нужно? можно убирать
	Debt            int64  `json:"debt"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

// Step ...
type Step int64

const (
	Step0None            Step = 0
	Step1Profession      Step = 1
	Step2Gender          Step = 2
	Step3World           Step = 3
	Step4Marriage        Step = 4
	Step5Children        Step = 5
	Step6Wishes          Step = 6
	Step7Turn            Step = 7
	Step8SalaryExtraName Step = 8
	Step9SalaryExtra     Step = 9
	Step10CostHouse      Step = 10
	Step11CostFood       Step = 11
	Step12CostTransport  Step = 12
	Step13CostCloth      Step = 13
	Step14CostExtraName  Step = 14
	Step15CostExtra      Step = 15
	Step16TotalIncome    Step = 16
	Step17TotalOutcome   Step = 17
	Step18Flow           Step = 18
	Step19Cash           Step = 19
)

// Gender ...
type Gender int64

const (
	GenderUnknown Gender = 0
	GenderMan     Gender = 1
	GenderWoman   Gender = 2
)

// World ...
type World int64

const (
	WorldUnknown World = 0
	WorldPoor    World = 1
	WorldMiddle  World = 2
	WorldRich    World = 3
)

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
func (m *MainData) GetStep() Step {
	if m == nil {
		return Step0None // zero-значение для Step
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
func (m *MainData) GetGender() Gender {
	if m == nil {
		return GenderUnknown
	}
	return m.Gender
}

// GetWorld ...
func (m *MainData) GetWorld() World {
	if m == nil {
		return WorldUnknown
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
