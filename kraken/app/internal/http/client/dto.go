package client

type BtcUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XXBTZUSD struct {
			Asks [][]interface{} `json:"asks"`
			Bids [][]interface{} `json:"bids"`
		} `json:"XXBTZUSD"`
	} `json:"result"`
}

type EthUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XXBTZUSD struct {
			Asks [][]interface{} `json:"asks"`
			Bids [][]interface{} `json:"bids"`
		} `json:"XETHZUSD"`
	} `json:"result"`
}

type TetherUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XXBTZUSD struct {
			Asks [][]interface{} `json:"asks"`
			Bids [][]interface{} `json:"bids"`
		} `json:"USDTZUSD"`
	} `json:"result"`
}

type XrpUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XXBTZUSD struct {
			Asks [][]interface{} `json:"asks"`
			Bids [][]interface{} `json:"bids"`
		} `json:"XXRPZUSD"`
	} `json:"result"`
}
