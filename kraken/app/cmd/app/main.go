package main

import (
	"fmt"
	"log"
	"time"

	"github.com/VrMolodyakov/crypto-comparing/kraken/internal/app"
	"github.com/VrMolodyakov/crypto-comparing/kraken/internal/http/client"
)

func main() {
	a := app.New()

	if err := a.ReadConfig(); err != nil {
		log.Fatal(err, "read config")
		return
	}

	a.InitLogger()
	api := client.New(a.Logger, 5*time.Second)
	dto, err := api.GetBtcRecentTrades()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dto)
}
