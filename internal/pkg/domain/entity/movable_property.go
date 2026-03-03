package entity

// MovableProperty ...
type MovableProperty struct {
	ID                  int64               `json:"id"`
	MainID              int64               `json:"main_id"`
	MovablePropertyType MovablePropertyType `json:"movable_property_type"`
	Name                string              `json:"name"` // TODO зачем нужно имя?
	Cost                int64               `json:"cost"`
	Term                int64               `json:"term"`
	Payment             int64               `json:"payment"`
}

// MovablePropertyType ...
type MovablePropertyType int64

const (
	MovablePropertyTypeUnknown  MovablePropertyType = 0
	MovablePropertyTypeAuto     MovablePropertyType = 1
	MovablePropertyTypeYacht    MovablePropertyType = 2
	MovablePropertyTypeAirplane MovablePropertyType = 3
)

// GetID ...
func (m *MovableProperty) GetID() int64 {
	if m == nil {
		return 0
	}
	return m.ID
}

// GetMainID ...
func (m *MovableProperty) GetMainID() int64 {
	if m == nil {
		return 0
	}
	return m.MainID
}

// GetMovablePropertyType ...
func (m *MovableProperty) GetMovablePropertyType() MovablePropertyType {
	if m == nil {
		return MovablePropertyTypeUnknown
	}
	return m.MovablePropertyType
}

// GetName ...
func (m *MovableProperty) GetName() string {
	if m == nil {
		return ""
	}
	return m.Name
}

// GetCost ...
func (m *MovableProperty) GetCost() int64 {
	if m == nil {
		return 0
	}
	return m.Cost
}

// GetTerm ...
func (m *MovableProperty) GetTerm() int64 {
	if m == nil {
		return 0
	}
	return m.Term
}

// GetPayment ...
func (m *MovableProperty) GetPayment() int64 {
	if m == nil {
		return 0
	}
	return m.Payment
}
