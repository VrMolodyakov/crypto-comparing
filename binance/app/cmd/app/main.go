package main

import (
	"fmt"
	"log"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/binance/internal/domain/currency/api"
)

func main() {
	client := api.New(5 * time.Second)
	trades, err := client.GetBtcRecentTrades(2)
	if err != nil {
		log.Fatal(err)
	}
	for _, trade := range trades {
		fmt.Println(trade)
	}
}
