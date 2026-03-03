package entity

// RealEstate ...
type RealEstate struct {
	ID             int64          `json:"id"`
	MainID         int64          `json:"main_id"`
	RealEstateType RealEstateType `json:"real_estate_type"`
	Name           string         `json:"name"` // TODO зачем нужно имя?
	Cost           int64          `json:"cost"`
	Term           int64          `json:"term"`
	Payment        int64          `json:"payment"`
}

// RealEstateType ...
type RealEstateType int64

const (
	RealEstateTypeUnknown RealEstateType = 0
	RealEstateTypeFlat    RealEstateType = 1
	RealEstateTypeLand    RealEstateType = 2
	RealEstateTypeChalet  RealEstateType = 3
	RealEstateTypeMansion RealEstateType = 4
)

// GetID ...
func (m *RealEstate) GetID() int64 {
	if m == nil {
		return 0
	}
	return m.ID
}

// GetMainID ...
func (m *RealEstate) GetMainID() int64 {
	if m == nil {
		return 0
	}
	return m.MainID
}

// GetRealEstateType ...
func (m *RealEstate) GetRealEstateType() RealEstateType {
	if m == nil {
		return RealEstateTypeUnknown
	}
	return m.RealEstateType
}

// GetName ...
func (m *RealEstate) GetName() string {
	if m == nil {
		return ""
	}
	return m.Name
}

// GetCost ...
func (m *RealEstate) GetCost() int64 {
	if m == nil {
		return 0
	}
	return m.Cost
}

// GetTerm ...
func (m *RealEstate) GetTerm() int64 {
	if m == nil {
		return 0
	}
	return m.Term
}

// GetPayment ...
func (m *RealEstate) GetPayment() int64 {
	if m == nil {
		return 0
	}
	return m.Payment
}
