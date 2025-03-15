package models

type PriceRequest struct {
	EventID string  `json:"eventId"`
	Outcome string  `json:"outcome"`
	Amount  float64 `json:"amount"`
}

type MarketState struct {
	YesShares float64 `json:"yesShares"`
	NoShares  float64 `json:"noShares"`
}

type PriceResponse struct {
	CurrentPrice   float64 `json:"currentPrice"`
	ImpactedPrice float64 `json:"impactedPrice"`
	PriceImpact   float64 `json:"priceImpact"`
	TotalCost     float64 `json:"totalCost"`
}
