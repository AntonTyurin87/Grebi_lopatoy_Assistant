package dto

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
)

const (
	BusinessTableName = "business"

	BusinessColumnID           = "id"
	BusinessColumnMainID       = "main_id"
	BusinessColumnBusinessType = "type"
	BusinessColumnName         = "name"
	BusinessColumnCost         = "cost"
	BusinessColumnProfit       = "profit"
)

var BusinessColumns = []string{
	BusinessColumnID,
	BusinessColumnMainID,
	BusinessColumnBusinessType,
	BusinessColumnName,
	BusinessColumnCost,
	BusinessColumnProfit,
}

// Business ...
type Business struct {
	ID           int64  `json:"id"`
	MainID       int64  `json:"main_id"`
	BusinessType int64  `json:"type"`
	Name         string `json:"name"`
	Cost         int64  `json:"cost"`
	Profit       int64  `json:"profit"`
}

// GetID ...
func (b *Business) GetID() int64 {
	if b == nil {
		return 0
	}
	return b.ID
}

// GetMainID ...
func (b *Business) GetMainID() int64 {
	if b == nil {
		return 0
	}
	return b.MainID
}

// GetBusinessType ...
func (b *Business) GetBusinessType() int64 {
	if b == nil {
		return 0
	}
	return b.BusinessType
}

// GetName ...
func (b *Business) GetName() string {
	if b == nil {
		return ""
	}
	return b.Name
}

// GetCost ...
func (b *Business) GetCost() int64 {
	if b == nil {
		return 0
	}
	return b.Cost
}

// GetProfit ...
func (b *Business) GetProfit() int64 {
	if b == nil {
		return 0
	}
	return b.Profit
}

// Entity ...
func (b *Business) Entity() *entity.Business {
	if b == nil {
		return nil
	}

	return &entity.Business{
		ID:           b.GetID(),
		MainID:       b.GetMainID(),
		BusinessType: entity.BusinessType(b.GetBusinessType()),
		Name:         b.GetName(),
		Cost:         b.GetCost(),
		Profit:       b.GetProfit(),
	}
}

// Businesses ...
type Businesses []*Business

// Entity ...
func (b Businesses) Entity() []*entity.Business {
	result := make([]*entity.Business, 0, len(b))
	for _, business := range b {
		if business != nil {
			result = append(result, business.Entity())
		}
	}
	return result
}

// BusinessDtoFromEntity ...
func BusinessDtoFromEntity(e *entity.Business) *Business {
	if e == nil {
		return nil
	}
	return &Business{
		ID:           e.GetID(),
		MainID:       e.GetMainID(),
		BusinessType: int64(e.GetBusinessType()),
		Name:         e.GetName(),
		Cost:         e.GetCost(),
		Profit:       e.GetProfit(),
	}
}

// BusinessDtoFromEntitySlice ...
func BusinessDtoFromEntitySlice(entities []*entity.Business) Businesses {
	if entities == nil {
		return nil
	}
	result := make(Businesses, 0, len(entities))
	for _, e := range entities {
		result = append(result, BusinessDtoFromEntity(e))
	}
	return result
}
