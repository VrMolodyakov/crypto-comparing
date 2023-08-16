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
		XXBTZUSD [][]interface{} `json:"XXBTZUSD"`
	} `json:"result"`
}

type EthUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XETHZUSD [][]interface{} `json:"XETHZUSD"`
	} `json:"result"`
}

type TetherUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		USDTZUSD [][]interface{} `json:"USDTZUSD"`
	} `json:"result"`
}

type XrpUsdTrades struct {
	Error  []interface{} `json:"error"`
	Result struct {
		XXRPZUSD [][]interface{} `json:"XXRPZUSD"`
	} `json:"result"`
}

func (t *BtcUsdTrades) Validate() error {
	count := len(t.Result.XXBTZUSD)
	if count == 0 {
		return errors.New("empty data")
	}
	for _, tradeInfo := range t.Result.XXBTZUSD {
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
		Volume:    volume,
		Timestamp: time.Unix(int64(trade[2].(float64)), 0),
		Buy:       trade[3].(string) == BUY,
		Sell:      trade[3].(string) == SELL,
	}, nil

}
