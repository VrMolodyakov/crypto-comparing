package api

import (
	"errors"
	"strconv"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/bybit/internal/domain/currency/model"
)

const (
	BUY  = "Buy"
	SELL = "Sell"
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

func (t *TradeContent) ConvertToInfo() (model.TradeInfo, error) {
	price, err := strconv.ParseFloat(t.Price, 64)
	if err != nil {
		return model.TradeInfo{}, err
	}
	volume, err := strconv.ParseFloat(t.Size, 64)
	if err != nil {
		return model.TradeInfo{}, err
	}
	timestamp, err := strconv.Atoi(t.Time)
	if err != nil {
		return model.TradeInfo{}, err
	}
	return model.TradeInfo{
		Price:     price,
		Size:      volume,
		Buy:       t.Side == BUY,
		Sell:      t.Side == SELL,
		Timestamp: time.UnixMilli(int64(timestamp)),
	}, nil
}
