package entity

// FinancialAsset ...
type FinancialAsset struct {
	ID                 int64              `json:"id"`
	MainID             int64              `json:"main_id"`
	FinancialAssetType FinancialAssetType `json:"financial_asset_type"`
	Ticker             string             `json:"ticker"` // TODO зачем нужно?
	Cost               int64              `json:"cost"`
	Quantity           int64              `json:"quantity"`
	Coupon             int64              `json:"coupon"`
	Time               int64              `json:"time"`
	Price              int64              `json:"price"`
	Rate               int64              `json:"rate"`
}

// FinancialAssetType ...
type FinancialAssetType int64

const (
	FinancialAssetsTypeUnknown FinancialAssetType = 0
	FinancialAssetsTypeStock   FinancialAssetType = 1
	FinancialAssetsTypeBound   FinancialAssetType = 2
	FinancialAssetsTypeDeposit FinancialAssetType = 3
)

// GetID ...
func (f *FinancialAsset) GetID() int64 {
	if f == nil {
		return 0
	}
	return f.ID
}

// GetMainID ...
func (f *FinancialAsset) GetMainID() int64 {
	if f == nil {
		return 0
	}
	return f.MainID
}

// GetFinancialAssetsType ...
func (f *FinancialAsset) GetFinancialAssetsType() FinancialAssetType {
	if f == nil {
		return FinancialAssetsTypeUnknown
	}
	return f.FinancialAssetType
}

// GetTicker ...
func (f *FinancialAsset) GetTicker() string {
	if f == nil {
		return ""
	}
	return f.Ticker
}

// GetCost ...
func (f *FinancialAsset) GetCost() int64 {
	if f == nil {
		return 0
	}
	return f.Cost
}

// GetQuantity ...
func (f *FinancialAsset) GetQuantity() int64 {
	if f == nil {
		return 0
	}
	return f.Quantity
}

// GetCoupon ...
func (f *FinancialAsset) GetCoupon() int64 {
	if f == nil {
		return 0
	}
	return f.Coupon
}

// GetTime ...
func (f *FinancialAsset) GetTime() int64 {
	if f == nil {
		return 0
	}
	return f.Time
}

// GetPrice ...
func (f *FinancialAsset) GetPrice() int64 {
	if f == nil {
		return 0
	}
	return f.Price
}

// GetRate ...
func (f *FinancialAsset) GetRate() int64 {
	if f == nil {
		return 0
	}
	return f.Rate
}
