package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/bybit/internal/domain/currency/model"
)

const (
	Bitcoin  string = "BTCUSDT"
	Ethereum string = "ETHUSDT"
	Doge     string = "DOGEUSDT"
	Xrp      string = "XRPUSDT"
	baseURL  string = "https://api.bybit.com/v5/market/recent-trade?category=spot&symbol=%s&limit=%d"
)

type apiClient struct {
	client http.Client
}

func New(timeout time.Duration) *apiClient {
	return &apiClient{
		client: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *apiClient) GetBtcRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Bitcoin, count)
	return c.getTradeInfo(url)
}

func (c *apiClient) GetEthRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Ethereum, count)
	return c.getTradeInfo(url)
}

func (c *apiClient) GetDogeRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Doge, count)
	return c.getTradeInfo(url)
}

func (c *apiClient) GetXrpRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Xrp, count)
	return c.getTradeInfo(url)
}

func (c *apiClient) getTradeInfo(url string) ([]model.TradeInfo, error) {
	response, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body due to %w", err)
	}
	var dto Trades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error = %w", err)
	}
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	trades := make([]model.TradeInfo, len(dto.Result.List))
	for i := range dto.Result.List {
		tradeInfo, err := dto.Result.List[i].ConvertToInfo()
		if err != nil {
			return nil, err
		}
		trades[i] = tradeInfo

	}
	return trades, nil
}
