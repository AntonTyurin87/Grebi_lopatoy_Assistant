package dto

import (
	"Grebi_lopatoy_Assistant/internal/pkg/domain/entity"
)

const (
	FinancialAssetsTableName = "financial_assets"

	FinancialAssetsColumnID                  = "id"
	FinancialAssetsColumnMainID              = "main_id"
	FinancialAssetsColumnFinancialAssetsType = "type"
	FinancialAssetsColumnTicker              = "ticker"
	FinancialAssetsColumnCost                = "cost"
	FinancialAssetsColumnQuantity            = "quantity"
	FinancialAssetsColumnCoupon              = "coupon"
	FinancialAssetsColumnTime                = "time"
	FinancialAssetsColumnPrice               = "price"
	FinancialAssetsColumnRate                = "rate"
)

var FinancialAssetsColumns = []string{
	FinancialAssetsColumnID,
	FinancialAssetsColumnMainID,
	FinancialAssetsColumnFinancialAssetsType,
	FinancialAssetsColumnTicker,
	FinancialAssetsColumnCost,
	FinancialAssetsColumnQuantity,
	FinancialAssetsColumnCoupon,
	FinancialAssetsColumnTime,
	FinancialAssetsColumnPrice,
	FinancialAssetsColumnRate,
}

// FinancialAssetsType - тип из entity
type FinancialAssetsType = entity.FinancialAssetType

// FinancialAssets ...
type FinancialAssets struct {
	ID                  int64  `json:"id"`
	MainID              int64  `json:"main_id"`
	FinancialAssetsType int64  `json:"financial_assets_type"`
	Ticker              string `json:"ticker"`
	Cost                int64  `json:"cost"`
	Quantity            int64  `json:"quantity"`
	Coupon              int64  `json:"coupon"`
	Time                int64  `json:"time"`
	Price               int64  `json:"price"`
	Rate                int64  `json:"rate"`
}

// GetID ...
func (f *FinancialAssets) GetID() int64 {
	if f == nil {
		return 0
	}
	return f.ID
}

// GetMainID ...
func (f *FinancialAssets) GetMainID() int64 {
	if f == nil {
		return 0
	}
	return f.MainID
}

// GetFinancialAssetsType ...
func (f *FinancialAssets) GetFinancialAssetsType() int64 {
	if f == nil {
		return 0
	}
	return f.FinancialAssetsType
}

// GetTicker ...
func (f *FinancialAssets) GetTicker() string {
	if f == nil {
		return ""
	}
	return f.Ticker
}

// GetCost ...
func (f *FinancialAssets) GetCost() int64 {
	if f == nil {
		return 0
	}
	return f.Cost
}

// GetQuantity ...
func (f *FinancialAssets) GetQuantity() int64 {
	if f == nil {
		return 0
	}
	return f.Quantity
}

// GetCoupon ...
func (f *FinancialAssets) GetCoupon() int64 {
	if f == nil {
		return 0
	}
	return f.Coupon
}

// GetTime ...
func (f *FinancialAssets) GetTime() int64 {
	if f == nil {
		return 0
	}
	return f.Time
}

// GetPrice ...
func (f *FinancialAssets) GetPrice() int64 {
	if f == nil {
		return 0
	}
	return f.Price
}

// GetRate ...
func (f *FinancialAssets) GetRate() int64 {
	if f == nil {
		return 0
	}
	return f.Rate
}

// Entity ...
func (f *FinancialAssets) Entity() *entity.FinancialAsset {
	if f == nil {
		return nil
	}

	return &entity.FinancialAsset{
		ID:                 f.GetID(),
		MainID:             f.GetMainID(),
		FinancialAssetType: entity.FinancialAssetType(f.GetFinancialAssetsType()),
		Ticker:             f.GetTicker(),
		Cost:               f.GetCost(),
		Quantity:           f.GetQuantity(),
		Coupon:             f.GetCoupon(),
		Time:               f.GetTime(),
		Price:              f.GetPrice(),
		Rate:               f.GetRate(),
	}
}

// FinancialAssetsSlice ...
type FinancialAssetsSlice []*FinancialAssets

// Entity ...
func (f FinancialAssetsSlice) Entity() []*entity.FinancialAsset {
	result := make([]*entity.FinancialAsset, 0, len(f))
	for _, asset := range f {
		if asset != nil {
			result = append(result, asset.Entity())
		}
	}
	return result
}

// FinancialAssetsDtoFromEntity ...
func FinancialAssetsDtoFromEntity(e *entity.FinancialAsset) *FinancialAssets {
	if e == nil {
		return nil
	}
	return &FinancialAssets{
		ID:                  e.GetID(),
		MainID:              e.GetMainID(),
		FinancialAssetsType: int64(e.GetFinancialAssetsType()),
		Ticker:              e.GetTicker(),
		Cost:                e.GetCost(),
		Quantity:            e.GetQuantity(),
		Coupon:              e.GetCoupon(),
		Time:                e.GetTime(),
		Price:               e.GetPrice(),
		Rate:                e.GetRate(),
	}
}

// FinancialAssetsDtoFromEntitySlice ...
func FinancialAssetsDtoFromEntitySlice(entities []*entity.FinancialAsset) FinancialAssetsSlice {
	if entities == nil {
		return nil
	}
	result := make(FinancialAssetsSlice, 0, len(entities))
	for _, e := range entities {
		result = append(result, FinancialAssetsDtoFromEntity(e))
	}
	return result
}
