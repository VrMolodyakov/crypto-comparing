package api

import (
	"errors"
	"strconv"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/kraken/internal/domain/currency/model"
)

const (
	apiResponseContentLength int = 7
	BUY                          = "b"
	SELL                         = "s"
	MARKET                       = "m"
	LIMIT                        = "l"
)

type BtcUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XBTUSDT [][]interface{} `json:"XBTUSDT"`
	} `json:"result"`
}

type EthUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		ETHUSDT [][]interface{} `json:"ETHUSDT"`
	} `json:"result"`
}

type TetherUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XDGUSDT [][]interface{} `json:"XDGUSDT"`
	} `json:"result"`
}

type XrpUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XRPUSDT [][]interface{} `json:"XRPUSDT"`
	} `json:"result"`
}

func (t *BtcUsdTrades) Validate() error {
	count := len(t.Result.XBTUSDT)
	if count == 0 {
		return errors.New("empty data")
	}
	for _, tradeInfo := range t.Result.XBTUSDT {
		if n := len(tradeInfo); n != apiResponseContentLength {
			return errors.New("incorrect response array length")
		}
	}
	return nil
}

func (t *EthUsdTrades) Validate() error {
	count := len(t.Result.ETHUSDT)
	if count == 0 {
		return errors.New("empty data")
	}
	for _, tradeInfo := range t.Result.ETHUSDT {
		if n := len(tradeInfo); n != apiResponseContentLength {
			return errors.New("incorrect response array length")
		}
	}
	return nil
}

func (t *TetherUsdTrades) Validate() error {
	count := len(t.Result.XDGUSDT)
	if count == 0 {
		return errors.New("empty data")
	}
	for _, tradeInfo := range t.Result.XDGUSDT {
		if n := len(tradeInfo); n != apiResponseContentLength {
			return errors.New("incorrect response array length")
		}
	}
	return nil
}

func (t *XrpUsdTrades) Validate() error {
	count := len(t.Result.XRPUSDT)
	if count == 0 {
		return errors.New("empty data")
	}
	for _, tradeInfo := range t.Result.XRPUSDT {
		if n := len(tradeInfo); n != apiResponseContentLength {
			return errors.New("incorrect response array length")
		}
	}
	return nil
}

func ConvertToInfo(trade []any) (model.TradeInfo, error) {
	priceString := trade[0].(string)
	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		return model.TradeInfo{}, err
	}
	volumeString := trade[1].(string)
	volume, err := strconv.ParseFloat(volumeString, 64)
	if err != nil {
		return model.TradeInfo{}, err
	}

	return model.TradeInfo{
		Price:     price,
		Size:      volume,
		Timestamp: time.Unix(int64(trade[2].(float64)), 0),
		Buy:       trade[3].(string) == BUY,
		Sell:      trade[3].(string) == SELL,
	}, nil

}
