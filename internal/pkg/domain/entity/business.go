package entity

// Business ...
type Business struct {
	ID           int64        `json:"id"`
	MainID       int64        `json:"main_id"`
	BusinessType BusinessType `json:"business_type"`
	Name         string       `json:"name"` // TODO зачем нужно имя?
	Cost         int64        `json:"cost"`
	Profit       int64        `json:"profit"`
}

// BusinessType ...
type BusinessType int64

const (
	BusinessTypeUnknown BusinessType = 0
	BusinessTypeSmall   BusinessType = 1
	BusinessTypeMedium  BusinessType = 2
	BusinessTypeBig     BusinessType = 3
)

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
func (b *Business) GetBusinessType() BusinessType {
	if b == nil {
		return BusinessTypeUnknown
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
