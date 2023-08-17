package api

import (
	"errors"
	"strconv"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/kucoin/internal/domain/currency/model"
)

const (
	okLowerCode = 200000
	okUpperCode = 260210
	BUY         = "buy"
	SELL        = "sell"
)

type Trades struct {
	Code string         `json:"code"`
	Data []TradeContent `json:"data"`
}

type TradeContent struct {
	Price string `json:"price"`
	Size  string `json:"size"`
	Side  string `json:"side"`
	Time  int64  `json:"time"`
}

func (t *Trades) Validate() error {
	code, err := strconv.Atoi(t.Code)
	if err != nil {
		return err
	}
	if code < okLowerCode || code >= okUpperCode {
		return errors.New(t.Code)
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
	return model.TradeInfo{
		Price:     price,
		Size:      volume,
		Buy:       t.Side == BUY,
		Sell:      t.Side == SELL,
		Timestamp: time.Unix(0, t.Time),
	}, nil
}
