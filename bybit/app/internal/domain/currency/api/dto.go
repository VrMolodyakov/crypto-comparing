package api

import "github.com/VrMolodyakov/crypto-comparing/bybit/internal/domain/currency/model"

type Trades struct {
	Result struct {
		Category string         `json:"category"`
		List     []TradeContent `json:"list"`
	} `json:"result"`
}

type TradeContent struct {
	Price string `json:"price"`
	Size  string `json:"size"`
	Side  string `json:"side"`
	Time  string `json:"time"`
}

func ConvertToInfo(trade TradeContent) (model.TradeInfo, error) {
	return model.TradeInfo{
		Price:  trade.Price,
		Volume: trade.Size,
	}
}
