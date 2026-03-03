package dto

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
)

const (
	MovablePropertyTableName = "movable_property"

	MovablePropertyColumnID                  = "id"
	MovablePropertyColumnMainID              = "main_id"
	MovablePropertyColumnMovablePropertyType = "type"
	MovablePropertyColumnName                = "name"
	MovablePropertyColumnCost                = "cost"
	MovablePropertyColumnTerm                = "term"
	MovablePropertyColumnPayment             = "payment"
)

var MovablePropertyColumns = []string{
	MovablePropertyColumnID,
	MovablePropertyColumnMainID,
	MovablePropertyColumnMovablePropertyType,
	MovablePropertyColumnName,
	MovablePropertyColumnCost,
	MovablePropertyColumnTerm,
	MovablePropertyColumnPayment,
}

// MovablePropertyType - тип из entity
type MovablePropertyType = entity.MovablePropertyType

// MovableProperty ...
type MovableProperty struct {
	ID                  int64  `json:"id"`
	MainID              int64  `json:"main_id"`
	MovablePropertyType int64  `json:"movable_property_type"`
	Name                string `json:"name"`
	Cost                int64  `json:"cost"`
	Term                int64  `json:"term"`
	Payment             int64  `json:"payment"`
}

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
func (m *MovableProperty) GetMovablePropertyType() int64 {
	if m == nil {
		return 0
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

// Entity ...
func (m *MovableProperty) Entity() *entity.MovableProperty {
	if m == nil {
		return nil
	}

	return &entity.MovableProperty{
		ID:                  m.GetID(),
		MainID:              m.GetMainID(),
		MovablePropertyType: entity.MovablePropertyType(m.GetMovablePropertyType()),
		Name:                m.GetName(),
		Cost:                m.GetCost(),
		Term:                m.GetTerm(),
		Payment:             m.GetPayment(),
	}
}

// MovableProperties ...
type MovableProperties []*MovableProperty

// Entity ...
func (m MovableProperties) Entity() []*entity.MovableProperty {
	result := make([]*entity.MovableProperty, 0, len(m))
	for _, prop := range m {
		if prop != nil {
			result = append(result, prop.Entity())
		}
	}
	return result
}

// MovablePropertyDtoFromEntity ...
func MovablePropertyDtoFromEntity(e *entity.MovableProperty) *MovableProperty {
	if e == nil {
		return nil
	}
	return &MovableProperty{
		ID:                  e.GetID(),
		MainID:              e.GetMainID(),
		MovablePropertyType: int64(e.GetMovablePropertyType()),
		Name:                e.GetName(),
		Cost:                e.GetCost(),
		Term:                e.GetTerm(),
		Payment:             e.GetPayment(),
	}
}

// MovablePropertyDtoFromEntitySlice ...
func MovablePropertyDtoFromEntitySlice(entities []*entity.MovableProperty) MovableProperties {
	if entities == nil {
		return nil
	}
	result := make(MovableProperties, 0, len(entities))
	for _, e := range entities {
		result = append(result, MovablePropertyDtoFromEntity(e))
	}
	return result
}
