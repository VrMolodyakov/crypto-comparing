package api

import (
	"errors"
	"strconv"

	"github.com/VrMolodyakov/crypto-comparing/bybit/internal/domain/currency/model"
)

const (
	BUY    = "Buy"
	SELL   = "Sell"
	MARKET = "m"
	LIMIT  = "l"
)

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

func (t *Trades) Validate() error {
	if len(t.Result.List) == 0 {
		return errors.New("trades is empty")
	}
	return nil
}

func (trade *TradeContent) ConvertToInfo() (model.TradeInfo, error) {
	price, err := strconv.ParseFloat(trade.Price, 64)
	if err != nil {
		return model.TradeInfo{}, err
	}
	volume, err := strconv.ParseFloat(trade.Size, 64)
	if err != nil {
		return model.TradeInfo{}, err
	}
	return model.TradeInfo{
		Price: price,
		Size:  volume,
		Buy:   trade.Side == BUY,
		Sell:  trade.Side == SELL,
	}, nil
}
