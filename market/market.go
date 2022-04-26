package market

type MarketOperation struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type TaxResponse struct {
	Tax float64 `json:"tax"`
}
