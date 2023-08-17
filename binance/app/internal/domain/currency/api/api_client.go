package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/binance/internal/domain/currency/model"
)

const (
	Bitcoin  string = "BTCUSDT"
	Ethereum string = "ETHUSDT"
	Doge     string = "DOGEUSDT"
	Xrp      string = "XRPUSDT"
	baseURL  string = "https://api.binance.com/api/v3/trades?symbol=%s&limit=%d"
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
	return c.getTradeInfo(url, "Bitcoin")
}

func (c *apiClient) GetEthRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Ethereum, count)
	return c.getTradeInfo(url, "Ethereum")
}

func (c *apiClient) GetDogeRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Doge, count)
	return c.getTradeInfo(url, "Doge")
}

func (c *apiClient) GetXrpRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Xrp, count)
	return c.getTradeInfo(url, "Xrp")
}

func (c *apiClient) getTradeInfo(url string, name string) ([]model.TradeInfo, error) {
	response, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body due to %w", err)
	}
	var tradesDto Trades
	err = json.Unmarshal(body, &tradesDto)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error = %w", err)
	}
	if err := tradesDto.Validate(); err != nil {
		return nil, err
	}
	trades := make([]model.TradeInfo, len(tradesDto))
	for i := range tradesDto {
		tradeInfo, err := tradesDto[i].ConvertToInfo()
		if err != nil {
			return nil, err
		}
		tradeInfo.Name = name
		trades[i] = tradeInfo

	}
	return trades, nil
}
