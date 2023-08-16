package model

import "time"

type TradeInfo struct {
	Name      string
	Price     float64
	Size      float64
	Timestamp time.Time
	Buy       bool
	Sell      bool
	Market    bool
	Limit     bool
}
