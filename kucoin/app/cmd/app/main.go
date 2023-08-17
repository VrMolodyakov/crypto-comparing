package main

import (
	"fmt"
	"log"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/kucoin/internal/domain/currency/api"
)

func main() {
	client := api.New(4 * time.Second)
	trades, err := client.GetXrpRecentTrades()
	if err != nil {
		log.Fatal(err)
	}
	for _, trade := range trades {
		fmt.Println(trade)
	}
}
