package service

import (
	"github.com/VrMolodyakov/crypto-comparing/kraken/internal/domain/currency/model"
	"github.com/VrMolodyakov/crypto-comparing/kraken/pkg/logging"
)

type CryptocurrencyClient interface {
	GetBtcRecentTrades(count int) ([]model.TradeInfo, error)
	GetEthRecentTrades(count int) ([]model.TradeInfo, error)
	GetDogeRecentTrades(count int) ([]model.TradeInfo, error)
	GetXrpRecentTrades(count int) ([]model.TradeInfo, error)
}

type service struct {
	logger logging.Logger
	client CryptocurrencyClient
}

func NewCryptocurrencyService(client CryptocurrencyClient, logger logging.Logger) *service {
	return &service{
		client: client,
		logger: logger,
	}
}

func (s *service) GetBtcTrades(count int) ([]model.TradeInfo, error) {
	s.logger.Debugf("try to get btc trades with count = %d", count)
	return s.client.GetBtcRecentTrades(count)
}
func (s *service) GetEthTrades(count int) ([]model.TradeInfo, error) {
	s.logger.Debugf("try to get ethereum trades with count = %d", count)
	return s.client.GetEthRecentTrades(count)
}
func (s *service) GetTetherTrades(count int) ([]model.TradeInfo, error) {
	s.logger.Debugf("try to get tether trades with count = %d", count)
	return s.client.GetDogeRecentTrades(count)
}
func (s *service) GetXrpTrades(count int) ([]model.TradeInfo, error) {
	s.logger.Debugf("try to get xrp trades with count = %d", count)
	return s.client.GetXrpRecentTrades(count)
}
