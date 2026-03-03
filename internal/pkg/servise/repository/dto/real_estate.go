package dto

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
)

const (
	RealEstateTableName = "real_estate"

	RealEstateColumnID             = "id"
	RealEstateColumnMainID         = "main_id"
	RealEstateColumnRealEstateType = "type"
	RealEstateColumnName           = "name"
	RealEstateColumnCost           = "cost"
	RealEstateColumnTerm           = "term"
	RealEstateColumnPayment        = "payment"
)

var RealEstateColumns = []string{
	RealEstateColumnID,
	RealEstateColumnMainID,
	RealEstateColumnRealEstateType,
	RealEstateColumnName,
	RealEstateColumnCost,
	RealEstateColumnTerm,
	RealEstateColumnPayment,
}

// RealEstateType - тип из entity
type RealEstateType = entity.RealEstateType

// RealEstate ...
type RealEstate struct {
	ID             int64  `json:"id"`
	MainID         int64  `json:"main_id"`
	RealEstateType int64  `json:"real_estate_type"`
	Name           string `json:"name"`
	Cost           int64  `json:"cost"`
	Term           int64  `json:"term"`
	Payment        int64  `json:"payment"`
}

// GetID ...
func (r *RealEstate) GetID() int64 {
	if r == nil {
		return 0
	}
	return r.ID
}

// GetMainID ...
func (r *RealEstate) GetMainID() int64 {
	if r == nil {
		return 0
	}
	return r.MainID
}

// GetRealEstateType ...
func (r *RealEstate) GetRealEstateType() int64 {
	if r == nil {
		return 0
	}
	return r.RealEstateType
}

// GetName ...
func (r *RealEstate) GetName() string {
	if r == nil {
		return ""
	}
	return r.Name
}

// GetCost ...
func (r *RealEstate) GetCost() int64 {
	if r == nil {
		return 0
	}
	return r.Cost
}

// GetTerm ...
func (r *RealEstate) GetTerm() int64 {
	if r == nil {
		return 0
	}
	return r.Term
}

// GetPayment ...
func (r *RealEstate) GetPayment() int64 {
	if r == nil {
		return 0
	}
	return r.Payment
}

// Entity ...
func (r *RealEstate) Entity() *entity.RealEstate {
	if r == nil {
		return nil
	}

	return &entity.RealEstate{
		ID:             r.GetID(),
		MainID:         r.GetMainID(),
		RealEstateType: entity.RealEstateType(r.GetRealEstateType()),
		Name:           r.GetName(),
		Cost:           r.GetCost(),
		Term:           r.GetTerm(),
		Payment:        r.GetPayment(),
	}
}

// RealEstates ...
type RealEstates []*RealEstate

// Entity ...
func (r RealEstates) Entity() []*entity.RealEstate {
	result := make([]*entity.RealEstate, 0, len(r))
	for _, estate := range r {
		if estate != nil {
			result = append(result, estate.Entity())
		}
	}
	return result
}

// RealEstateDtoFromEntity ...
func RealEstateDtoFromEntity(e *entity.RealEstate) *RealEstate {
	if e == nil {
		return nil
	}
	return &RealEstate{
		ID:             e.GetID(),
		MainID:         e.GetMainID(),
		RealEstateType: int64(e.GetRealEstateType()),
		Name:           e.GetName(),
		Cost:           e.GetCost(),
		Term:           e.GetTerm(),
		Payment:        e.GetPayment(),
	}
}

// RealEstateDtoFromEntitySlice ...
func RealEstateDtoFromEntitySlice(entities []*entity.RealEstate) RealEstates {
	if entities == nil {
		return nil
	}
	result := make(RealEstates, 0, len(entities))
	for _, e := range entities {
		result = append(result, RealEstateDtoFromEntity(e))
	}
	return result
}
