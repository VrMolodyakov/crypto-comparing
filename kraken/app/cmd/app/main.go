package main

import (
	"fmt"
	"log"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/kraken/internal/app"
	"github.com/VrMolodyakov/crypto-comparing/kraken/internal/domain/currency/api"
)

func main() {
	a := app.New()

	if err := a.ReadConfig(); err != nil {
		log.Fatal(err, "read config")
		return
	}

	a.InitLogger()
	apiClient := api.New(5 * time.Second)
	trades, err := apiClient.GetDogeRecentTrades(2)
	if err != nil {
		log.Fatal(err)
	}
	for _, trade := range trades {
		fmt.Println(trade)
	}

}
