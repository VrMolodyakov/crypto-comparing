package api

import (
	"errors"
	"strconv"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/binance/internal/domain/currency/model"
)

type Trades []Trade

type Trade struct {
	Price        string `json:"price"`
	Size         string `json:"qty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

func (t *Trades) Validate() error {
	if len(*t) == 0 {
		return errors.New("empty trade list")
	}
	return nil
}

func (t *Trade) ConvertToInfo() (model.TradeInfo, error) {
	price, err := strconv.ParseFloat(t.Price, 64)
	if err != nil {
		return model.TradeInfo{}, err
	}
	size, err := strconv.ParseFloat(t.Size, 64)
	if err != nil {
		return model.TradeInfo{}, err
	}
	return model.TradeInfo{
		Price:     price,
		Size:      size,
		Timestamp: time.UnixMilli(t.Time),
		Buy:       t.IsBuyerMaker,
		Sell:      !t.IsBuyerMaker,
	}, nil
}
