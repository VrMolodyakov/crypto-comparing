package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/kraken/pkg/logging"
)

const (
	Bitcoin  string = "XBTUSD"
	Ethereum string = "XETHZUSD"
	Tether   string = "USDTZUSD"
	Xrp      string = "XBTUSD"
)

type apiClient struct {
	logger logging.Logger
	client http.Client
}

func New(logger logging.Logger, timeout time.Duration) *apiClient {
	return &apiClient{
		client: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *apiClient) GetBtcRecentTrades() (BtcUsdTrades, error) {
	url := fmt.Sprintf("https://api.kraken.com/0/public/Trades?pair=%s", Bitcoin)
	response, err := c.client.Get(url)
	if err != nil {
		return BtcUsdTrades{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return BtcUsdTrades{}, fmt.Errorf("cannot read response body due to %w:", err)
	}

	var dto BtcUsdTrades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return BtcUsdTrades{}, fmt.Errorf("JSON unmarshal error = %w:", err)
	}
	return dto, nil
}

func (c *apiClient) GetEthRecentTrades() (EthUsdTrades, error) {
	url := fmt.Sprintf("https://api.kraken.com/0/public/Trades?pair=%s", Ethereum)
	response, err := c.client.Get(url)
	if err != nil {
		return EthUsdTrades{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return EthUsdTrades{}, fmt.Errorf("cannot read response body due to %w:", err)
	}

	var dto EthUsdTrades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return EthUsdTrades{}, fmt.Errorf("JSON unmarshal error = %w:", err)
	}
	return dto, nil
}

func (c *apiClient) GetTetherRecentTrades() (TetherUsdTrades, error) {
	url := fmt.Sprintf("https://api.kraken.com/0/public/Trades?pair=%s", Tether)
	response, err := c.client.Get(url)
	if err != nil {
		return TetherUsdTrades{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return TetherUsdTrades{}, fmt.Errorf("cannot read response body due to %w:", err)
	}

	var dto TetherUsdTrades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return TetherUsdTrades{}, fmt.Errorf("JSON unmarshal error = %w:", err)
	}
	return dto, nil
}

func (c *apiClient) GetXrpRecentTrades() (XrpUsdTrades, error) {
	url := fmt.Sprintf("https://api.kraken.com/0/public/Trades?pair=%s", Xrp)
	response, err := c.client.Get(url)
	if err != nil {
		return XrpUsdTrades{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return XrpUsdTrades{}, fmt.Errorf("cannot read response body due to %w:", err)
	}

	var dto XrpUsdTrades
	err = json.Unmarshal(body, &dto)
	if err != nil {
		return XrpUsdTrades{}, fmt.Errorf("JSON unmarshal error = %w:", err)
	}
	return dto, nil
}
