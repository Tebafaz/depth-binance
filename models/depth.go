package models

type Depth struct {
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type PriceQuantity struct {
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

type DepthResponse struct {
	Bids        []PriceQuantity `json:"bids"`
	Asks        []PriceQuantity `json:"asks"`
	BidQuantSum float64         `json:"sum_of_bid_quantity"`
	BidPriceSum float64         `json:"sum_of_bid_price"`
	AskQuantSum float64         `json:"sum_of_ask_quantity"`
	AskPriceSum float64         `json:"sum_of_ask_price"`
}
