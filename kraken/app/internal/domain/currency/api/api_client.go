package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/kraken/internal/domain/currency/model"
)

const (
	Bitcoin  string = "XBTUSD"
	Ethereum string = "XETHZUSD"
	Tether   string = "USDTZUSD"
	Xrp      string = "XXRPZUSD"
	baseURL  string = "https://api.kraken.com/0/public/Trades?pair=%s&count=%d"
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

	var dto BtcUsdTrades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error = %w", err)
	}
	if len(dto.Error) != 0 {
		return nil, errors.New("api rate limit exceeded")
	}
	var trades []model.TradeInfo
	for _, d := range dto.Result.XXBTZUSD {
		trade, err := ConvertToInfo(d)
		if err != nil {
			return nil, fmt.Errorf("cannot convert to model due to = %w", err)
		}
		trade.Name = "Bitcoin"
		trades = append(trades, trade)
	}
	return trades, nil
}

func (c *apiClient) GetEthRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Ethereum, count)
	response, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body due to %w", err)
	}

	var dto EthUsdTrades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshal error = %w", err)
	}
	if len(dto.Error) != 0 {
		return nil, errors.New("api rate limit exceeded")
	}
	var trades []model.TradeInfo
	for _, d := range dto.Result.XETHZUSD {
		trade, err := ConvertToInfo(d)
		if err != nil {
			return nil, fmt.Errorf("cannot convert to model due to = %w", err)
		}
		trade.Name = "Ethereum"
		trades = append(trades, trade)
	}
	return trades, nil
}

func (c *apiClient) GetTetherRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Tether, count)
	response, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body due to %w", err)
	}

	var dto TetherUsdTrades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshal error = %w", err)
	}
	if len(dto.Error) != 0 {
		return nil, errors.New("api rate limit exceeded")
	}
	var trades []model.TradeInfo
	for _, d := range dto.Result.USDTZUSD {
		trade, err := ConvertToInfo(d)
		if err != nil {
			return nil, fmt.Errorf("cannot convert to model due to = %w", err)
		}
		trade.Name = "Tether"
		trades = append(trades, trade)
	}
	return trades, nil
}

func (c *apiClient) GetXrpRecentTrades(count int) ([]model.TradeInfo, error) {
	url := fmt.Sprintf(baseURL, Xrp, count)
	response, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body due to %w", err)
	}

	var dto XrpUsdTrades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshal error = %w", err)
	}
	if len(dto.Error) != 0 {
		return nil, errors.New("api rate limit exceeded")
	}
	var trades []model.TradeInfo
	for _, d := range dto.Result.XXRPZUSD {
		trade, err := ConvertToInfo(d)
		if err != nil {
			return nil, fmt.Errorf("cannot convert to model due to = %w", err)
		}
		trade.Name = "Xrp"
		trades = append(trades, trade)
	}
	return trades, nil
}
