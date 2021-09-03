package models

// MarketResponse ...
type MarketResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  []ResultPayload `json:"result"`
}

// ResultPayload ...
type ResultPayload struct {
	MarketName string  `json:"MarketName"`
	High       float64 `json:"High"`
	Low        float64 `json:"Low"`
	Volume     float64 `json:"Volume"`
	Timestamp  string  `json:"Timestamp"`
}
