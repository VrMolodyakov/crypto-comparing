package v1

import tradePb "github.com/VrMolodyakov/crypto-comparing/aggregator/app/gen/go/proto/aggregator_service/trade/v1"

type server struct {
	tradePb.UnimplementedTradeServiceServer
}
