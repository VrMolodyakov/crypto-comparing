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
	Bitcoin  string = "XBTUSD"
	Ethereum string = "XETHZUSD"
	Tether   string = "USDTZUSD"
	Xrp      string = "XXRPZUSD"
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
	return nil, nil
}

func (c *apiClient) GetEthRecentTrades(count int) ([]model.TradeInfo, error) {
	return nil, nil
}

func (c *apiClient) GetTetherRecentTrades(count int) ([]model.TradeInfo, error) {
	return nil, nil
}

func (c *apiClient) GetXrpRecentTrades(count int) ([]model.TradeInfo, error) {
	return nil, nil
}
