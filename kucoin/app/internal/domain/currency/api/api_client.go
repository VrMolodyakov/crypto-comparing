package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/kucoin/internal/domain/currency/model"
)

const (
	Bitcoin  string = "BTC"
	USDT     string = "USDT"
	Ethereum string = "ETH"
	Doge     string = "DOGE"
	Xrp      string = "XRP"
	baseURL  string = "https://api.kucoin.com/api/v1/market/histories?symbol=%s-%s"
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

func (c *apiClient) GetBtcRecentTrades() ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Bitcoin, USDT)
	return c.getTradeInfo(url, "Bitcoin")
}

func (c *apiClient) GetEthRecentTrades() ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Ethereum, USDT)
	return c.getTradeInfo(url, "Ethereum")
}

func (c *apiClient) GetDogeRecentTrades() ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Doge, USDT)
	return c.getTradeInfo(url, "Doge")
}

func (c *apiClient) GetXrpRecentTrades() ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Xrp, USDT)
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
	var dto Trades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error = %w", err)
	}
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	trades := make([]model.TradeInfo, len(dto.Data))
	for i := range dto.Data {
		tradeInfo, err := dto.Data[i].ConvertToInfo()
		if err != nil {
			return nil, err
		}
		tradeInfo.Name = name
		trades[i] = tradeInfo

	}
	return trades, nil
}
