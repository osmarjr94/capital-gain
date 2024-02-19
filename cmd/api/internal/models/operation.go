package models

type Operation struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type TaxResult struct {
	Tax int `json:"tax"`
}
