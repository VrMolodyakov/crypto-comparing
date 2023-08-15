package client

type Order struct {
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	Timestamp int64  `json:"timestamp"`
}

type OrderBook struct {
	Asks []Order `json:"asks"`
	Bids []Order `json:"bids"`
}

type Result struct {
	OrderBook map[string]OrderBook `json:"result"`
}

type Response struct {
	Error  []string `json:"error"`
	Result Result   `json:"result"`
}
