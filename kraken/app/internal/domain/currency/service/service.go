package service

import (
	"github.com/VrMolodyakov/crypto-comparing/kraken/internal/domain/currency/model"
	"github.com/VrMolodyakov/crypto-comparing/kraken/pkg/logging"
)

type CurrencyClient interface {
	GetOrderBook()
}

type currencyInfo struct {
	logger logging.Logger
	client CurrencyClient
}

func New(client CurrencyClient, logger logging.Logger) *currencyInfo {
	return &currencyInfo{
		client: client,
		logger: logger,
	}
}

func (c *currencyInfo) GetOrderBook() []model.Order {
	return nil
}
